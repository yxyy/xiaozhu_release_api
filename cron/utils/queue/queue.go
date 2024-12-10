package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
	"time"
)

const batchCount = 10

type Processor interface {
	Run(*Queue, string) error
}

type BatchProcessor interface {
	Run(*Queue, []string) (redoMsg []string, err error)
}

type Queue struct {
	Ctx            context.Context
	processor      Processor
	batchProcessor BatchProcessor
	coupler        Coupler
	jobChan        chan struct{}
	Log            *log.Entry
	errCount       int
	maxErr         int
	queue          string
	failZSort      string
	maxRetries     int8
	reTryTime      []int
	reTryNow       bool
	ts             int
}

func NewQueue(queue string, processor Processor) *Queue {
	return &Queue{
		Ctx:        context.Background(),
		processor:  processor,
		coupler:    DefaultCoupler,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		queue:      queue,
		failZSort:  queue + "_failed",
		Log:        log.WithField("queue", queue),
		reTryTime:  []int{60, 300, 1800},
		ts:         1,
	}
}

func NewBatchQueue(queue string, processor BatchProcessor, ts int) *Queue {
	if ts < 1 {
		ts = batchCount
	}
	return &Queue{
		Ctx:            context.Background(),
		batchProcessor: processor,
		coupler:        DefaultCoupler,
		jobChan:        make(chan struct{}, 100),
		maxRetries:     3,
		maxErr:         5,
		queue:          queue,
		failZSort:      queue + "_failed",
		Log:            log.WithField("queue", queue),
		reTryTime:      []int{60, 300, 1800},
		ts:             ts,
	}
}

func NewQueueWithContext(ctx context.Context, queue string, processor Processor) *Queue {
	return &Queue{
		Ctx:        ctx,
		processor:  processor,
		coupler:    DefaultCoupler,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		queue:      queue,
		failZSort:  queue + "_failed",
		Log:        log.WithField("queue", queue),
		reTryTime:  []int{60, 300, 1800},
	}
}

func NewBatchQueueWithContext(ctx context.Context, queue string, processor BatchProcessor, ts int) *Queue {
	if ts < 1 {
		ts = batchCount
	}
	return &Queue{
		Ctx:            ctx,
		batchProcessor: processor,
		coupler:        DefaultCoupler,
		jobChan:        make(chan struct{}, 100),
		maxRetries:     3,
		maxErr:         5,
		queue:          queue,
		failZSort:      queue + "_failed",
		Log:            log.WithField("queue", queue),
		reTryTime:      []int{60, 300, 1800},
		ts:             ts,
	}
}

func (q *Queue) Run() {
	defer q.recover()
	q.init()
	go q.reDo()
	for {
		select {
		case <-q.Ctx.Done():
			return
		default:
			q.handleMaxError()
			msg, err := q.Pops()
			if err != nil {
				q.handleRedisError(err)
				continue
			}
			q.AddJob(msg)
			q.errCount = 0
		}
	}
}

func (q *Queue) init() {
	if q.coupler == nil {
		q.Log.Error("队列连接器未设置")
		return
	}

	if int(q.maxRetries) != len(q.reTryTime) {
		q.Log.Error("队列初始化失败，重试配置不一致")
		return
	}

	q.reTryNow = q.isImmediateRetry()

	if q.processor != nil && q.batchProcessor != nil {
		q.Log.Error("单次和批量接口不能同时存在")
		return
	}

	if q.ts < 1 {
		if q.processor != nil {
			q.ts = 1
		} else {
			q.ts = batchCount
		}
	}
}

func (q *Queue) SetRetry(maxRetries int8, reTryTime []int) error {
	// 验证 maxRetries 的有效性
	if maxRetries < 0 || maxRetries > 5 {
		return fmt.Errorf("maxRetries should be between 0 and 5, got %d", maxRetries)
	}

	// 验证 reTryTime 数组长度
	if len(reTryTime) != int(maxRetries) {
		return fmt.Errorf("reTryTime array length must equal maxRetries, expected %d, got %d", maxRetries, len(reTryTime))
	}

	// 验证 reTryTime 中的负值
	for _, v := range reTryTime {
		if v < 0 {
			return fmt.Errorf("reTryTime cannot contain negative values, got %v", reTryTime)
		}
	}

	// 排序 reTryTime 数组（如不需要可以删除这行）
	sort.Ints(reTryTime)

	q.maxRetries = maxRetries
	q.reTryTime = reTryTime

	return nil
}

func (q *Queue) Pops() ([]string, error) {
	if q.processor != nil {
		return q.coupler.Pop(q.Ctx, q.queue)
	}

	return q.coupler.BatchPop(q.Ctx, q.queue, q.ts)
}

func (q *Queue) AddJob(msg []string) {
	defer q.jobRecover(msg)
	q.jobChan <- struct{}{}
	go func() {
		defer q.jobDone()
		if q.processor != nil {
			if err := q.processor.Run(q, msg[0]); err != nil {
				q.Log.Errorf("队列处理有误:%s，准备重新入队...", err)
				// 类型断言判断是否实现了 Retry 方法
				if retryProcessor, ok := q.processor.(interface{ Retry(*Queue, string) }); ok {
					retryProcessor.Retry(q, msg[0]) // 调用 processor 自己的 Retry 方法
				} else {
					q.Retry(msg[0])
				}
			}
		}

		if q.batchProcessor != nil {
			redoMsg, err := q.batchProcessor.Run(q, msg) // 返回重试消息的要自己为何re_try
			if err != nil {
				q.Log.Errorf("队列处理有误:%s，准备重新入队...", err)
				if len(redoMsg) > 0 {
					// 类型断言判断是否实现了 Retry 方法
					if retryProcessor, ok := q.processor.(interface{ Retry(*Queue, []string) }); ok {
						retryProcessor.Retry(q, redoMsg) // 调用 processor 自己的 Retry 方法
					} else {
						for _, v := range redoMsg {
							q.Retry(v) // 调用通用的 Retry 方法
						}
					}
				}
			}
		}

	}()
}

func (q *Queue) jobDone() {
	<-q.jobChan
}

func (q *Queue) handleRedisError(err error) {
	if errors.Is(err, redis.Nil) {
		q.Log.Info("队列暂无数据，等待中...")
		time.Sleep(5 * time.Second)
	} else {
		q.Log.Errorf("Redis 错误: %v", err)
		q.errCount++
	}
}

func (q *Queue) handleMaxError() {
	if q.errCount > q.maxErr {
		q.Log.Warnf("连续错误超过最大次数，休眠 5 分钟")
		time.Sleep(5 * 60 * time.Second)
		q.errCount = 0 // 清零错误计数，避免死循环
	}
}

// Retry 自定义实现Retry(q *Queue)，就调用自己的Retry
func (q *Queue) Retry(msg string) {
	if msg == "" {
		return
	}

	var mpData = make(map[string]any)
	err := json.Unmarshal([]byte(msg), &mpData)
	if err != nil {
		q.Log.Errorf("重试任务反序列化失败: %v", err)
		return
	}

	reTry, ok := mpData["re_try"].(float64)
	if !ok {
		reTry = 0
	}

	if int8(reTry) >= q.maxRetries {
		q.Log.Warn("达到最大重试次数，不再重试")
		return
	}

	reTry++
	mpData["re_try"] = reTry
	data, err := json.Marshal(mpData)
	if err != nil {
		q.Log.Errorf("重试任务序列化失败: %v", err)
		return
	}

	if q.reTryNow {
		err = q.coupler.Push(q.Ctx, q.queue, data)
	} else {
		err = q.coupler.FailAdd(q.Ctx, q.failZSort, float64(time.Now().Add(q.calculateDelay(int(reTry))).Unix()), data)
	}

	if err != nil {
		q.Log.Errorf("任务重新入队失败: %v", err)
	}
}

func (q *Queue) isImmediateRetry() bool {
	for _, v := range q.reTryTime {
		if v > 0 {
			return false
		}
	}

	return true
}

func (q *Queue) recover() {
	if err := recover(); err != nil {
		q.Log.Errorf("Recovered panic in queue %s: %v ", q.queue, err)
	}
}

func (q *Queue) jobRecover(msg []string) {
	if err := recover(); err != nil {
		q.Log.Errorf("jobRecover panic in queue %s: %v", q.queue, err)
		for _, v := range msg {
			q.Retry(v)
		}
	}
}

func (q *Queue) calculateDelay(reTry int) time.Duration {
	if len(q.reTryTime) < reTry {
		return time.Minute
	}

	return time.Duration(q.reTryTime[reTry-1]) * time.Second
}

func (q *Queue) reDo() {
	tickerInterval := time.Second * 10 // 默认的检查间隔时间
	ticker := time.NewTicker(tickerInterval)
	defer ticker.Stop()
	var lastTime int64
	pageSize := 30

	for {
		select {
		case <-q.Ctx.Done():
			return
		case <-ticker.C:
			now := time.Now().Unix()
			minTime := fmt.Sprintf("%d", lastTime)
			maxTime := fmt.Sprintf("%d", now)
			fmt.Printf("正在检查%s ~ %s\n", minTime, maxTime)
			// 获取失败队列中的任务数量
			count, err := q.coupler.FailNum(q.Ctx, q.failZSort, minTime, maxTime)
			if err != nil {
				q.Log.Errorf("获取失败队列任务数量失败： %v", err)
				continue
			}

			page := int(math.Ceil(float64(count) / float64(pageSize)))
			for i := 0; i < page; i++ {
				offset := i * pageSize
				result, err := q.coupler.FailRangeByScore(q.Ctx, q.failZSort, minTime, maxTime, int64(offset), int64(pageSize))
				if err != nil {
					q.Log.Errorf("获取分数集合数据失败： %s", err)
					break
				}
				if err = q.coupler.Push(q.Ctx, q.queue, result); err != nil {
					q.Log.Errorf("从新入队列失败： %s", err)
					break
				}
			}

			err = q.coupler.FailRemRangeByScore(q.Ctx, q.failZSort, minTime, maxTime)
			if err != nil {
				q.Log.Errorf("移除重新入队元素失败： %s", err)
			}

			lastTime = now
			// 调整 ticker 的间隔时间
			interval := q.tickerInterval(count)
			if interval > 0 && time.Duration(interval) != tickerInterval {
				ticker.Reset(time.Duration(interval) * time.Second)
				tickerInterval = time.Duration(interval) // 缓存当前间隔时间，避免频繁重置
			}
		}
	}
}

func (q *Queue) tickerInterval(count int64) int {
	var tickerInterval int

	switch {
	case count < 10:
		tickerInterval = 30 // 任务较少时，每30秒检查一次
	case count >= 10 && count <= 50:
		tickerInterval = 15 // 任务适中时，每15秒检查一次
	case count > 100:
		tickerInterval = 5 // 任务较多时，每5秒检查一次
	default:
		tickerInterval = 10 // 默认间隔
	}

	return tickerInterval
}

func (q *Queue) Len() (int64, error) {
	return q.coupler.Len(q.Ctx, q.queue)
}
