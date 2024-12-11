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
	"sync"
	"time"
)

const batchCount = 10

var redoOnce = sync.Map{}

type Processor interface {
	Run(*Queue, string) error
}

type BatchProcessor interface {
	Run(*Queue, []string) (redoMsg []string, err error)
}

type Queue struct {
	Ctx            context.Context
	name           string
	failName       string
	errCount       int
	maxErr         int
	maxRetries     int8
	reTryTime      []int
	reTryNow       bool
	ts             int
	coupler        Coupler
	processor      Processor
	batchProcessor BatchProcessor
	jobChan        chan struct{}
	Log            *log.Entry
	isTest         int
}

func NewQueue(queue string, processor Processor) *Queue {
	return &Queue{
		Ctx:        context.Background(),
		processor:  processor,
		coupler:    DefaultCoupler,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		name:       queue,
		failName:   queue + "_failed",
		Log:        log.WithField("name", queue),
		reTryTime:  []int{60, 300, 1800},
		ts:         1,
		isTest:     1,
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
		name:           queue,
		failName:       queue + "_failed",
		Log:            log.WithField("name", queue),
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
		name:       queue,
		failName:   queue + "_failed",
		Log:        log.WithField("name", queue),
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
		name:           queue,
		failName:       queue + "_failed",
		Log:            log.WithField("name", queue),
		reTryTime:      []int{60, 300, 1800},
		ts:             ts,
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

func (q *Queue) Run() {
	defer q.recover()
	if err := q.init(); err != nil {
		fmt.Println(err)
		q.Log.Error(err)
		return
	}
	if q.isTest >= 1 {
		return
	}
	for {
		select {
		case <-q.Ctx.Done():
			return
		default:
			q.handleSleep()
			msg, err := q.Pops()
			if err != nil {
				q.handleError(err)
				continue
			}
			q.AddJob(msg)
			q.errCount = 0
		}
	}
}

func (q *Queue) init() error {
	if q.coupler == nil {
		return fmt.Errorf("队列连接器未设置")
	}

	if int(q.maxRetries) != len(q.reTryTime) {
		return fmt.Errorf("队列初始化失败，重试配置不一致")
	}
	q.reTryNow = q.isImmediateRetry()

	if q.processor != nil && q.batchProcessor != nil {
		return fmt.Errorf("单次和批量接口不能同时存在")
	}

	if q.ts < 1 {
		if q.processor != nil {
			q.ts = 1
		} else {
			q.ts = batchCount
		}
	}

	RegisterMonitor(q.name, q)

	q.registerRoDo()

	return nil
}

func (q *Queue) Pops() ([]string, error) {
	if q.processor != nil {
		return q.coupler.Pop(q.Ctx, q.name)
	}

	return q.coupler.BatchPop(q.Ctx, q.name, q.ts)
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
		err = q.coupler.Push(q.Ctx, q.name, data)
	} else {
		err = q.coupler.FailAdd(q.Ctx, q.failName, q.nextReTryTime(int(reTry)), data)
	}

	if err != nil {
		q.Log.Errorf("任务重新入队失败: %v", err)
	}
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
			prev := fmt.Sprintf("%d", lastTime)
			next := fmt.Sprintf("%d", now)
			fmt.Printf("正在检查%s ~ %s\n", prev, next)
			// 获取失败队列中的任务数量
			count, err := q.coupler.FailNum(q.Ctx, q.failName, prev, next)
			if err != nil {
				q.Log.Errorf("获取失败队列任务数量失败： %v", err)
				continue
			}

			page := int(math.Ceil(float64(count) / float64(pageSize)))
			for i := 0; i < page; i++ {
				offset := i * pageSize
				result, err := q.coupler.FailRangeByScore(q.Ctx, q.failName, prev, next, int64(offset), int64(pageSize))
				if err != nil {
					q.Log.Errorf("获取分数集合数据失败： %s", err)
					break
				}
				if err = q.coupler.Push(q.Ctx, q.name, result); err != nil {
					q.Log.Errorf("从新入队列失败： %s", err)
					break
				}
			}

			err = q.coupler.FailRemRangeByScore(q.Ctx, q.failName, prev, next)
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

func (q *Queue) Len() (int64, error) {
	return q.coupler.Len(q.Ctx, q.name)
}

func (q *Queue) jobDone() {
	<-q.jobChan
}

func (q *Queue) handleError(err error) {
	if errors.Is(err, redis.Nil) {
		q.Log.Info("队列暂无数据，等待中...")
		time.Sleep(5 * time.Second)
	} else {
		q.Log.Errorf("Redis 错误: %v", err)
		q.errCount++
	}
}

func (q *Queue) handleSleep() {
	if q.errCount > q.maxErr {
		q.Log.Warnf("连续错误超过最大次数，休眠 5 分钟")
		time.Sleep(5 * 60 * time.Second)
		q.errCount = 0 // 清零错误计数，避免死循环
	}
}

func (q *Queue) recover() {
	if err := recover(); err != nil {
		q.Log.Errorf("Recovered panic in name %s: %v ", q.name, err)
	}
}

func (q *Queue) jobRecover(msg []string) {
	if err := recover(); err != nil {
		q.Log.Errorf("jobRecover panic in name %s: %v", q.name, err)
		for _, v := range msg {
			q.Retry(v)
		}
	}
}

func (q *Queue) nextReTryTime(reTry int) float64 {
	var duration time.Duration
	if len(q.reTryTime) < reTry {
		duration = time.Minute
	} else {
		duration = time.Duration(q.reTryTime[reTry-1]) * time.Second
	}

	return float64(time.Now().Add(duration).Unix())
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

func (q *Queue) isImmediateRetry() bool {
	for _, v := range q.reTryTime {
		if v > 0 {
			return false
		}
	}

	return true
}

func (q *Queue) registerRoDo() {
	if q.isTest >= 1 {
		return
	}
	var onceRedo *sync.Once
	if val, ok := redoOnce.Load(q.name); ok {
		if tmp, ok := val.(*sync.Once); ok {
			onceRedo = tmp
		}
	}
	// 如果 onceRedo 为空，创建一个新的
	if onceRedo == nil {
		onceRedo = &sync.Once{}
		redoOnce.Store(q.name, onceRedo)
	}

	go onceRedo.Do(q.reDo)

}
