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
	"xiaozhu/utils"
)

type BatchProcessor interface {
	Run(*BatchQueue, []string) error
}

type BatchQueue struct {
	ctx        context.Context
	processor  BatchProcessor
	jobChan    chan struct{}
	log        *log.Entry
	errCount   int
	maxErr     int
	queue      string
	failZSort  string
	Data       string
	maxRetries int8
	reTryTime  []int
	reTryNow   bool
	ts         int
}

func NewBatchQueue(queue string, processor BatchProcessor, ts int) *BatchQueue {
	if ts < 1 {
		ts = 10
	}
	return &BatchQueue{
		ctx:        context.Background(),
		processor:  processor,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		queue:      queue,
		failZSort:  queue + "_failed",
		log:        log.WithField("queue", queue),
		reTryTime:  []int{60, 300, 1800},
		ts:         ts,
	}
}

func NewBatchQueueWithContext(ctx context.Context, queue string, processor BatchProcessor) *BatchQueue {
	return &BatchQueue{
		ctx:        ctx,
		processor:  processor,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		queue:      queue,
		failZSort:  queue + "_failed",
		log:        log.WithField("queue", queue),
		reTryTime:  []int{60, 300, 1800},
		ts:         10,
	}
}

func NewBatchRetryQueue(queue string, processor BatchProcessor, maxRetries int8, reTryTime []int) (*BatchQueue, error) {
	// 验证 maxRetries 的有效性
	if maxRetries < 0 || maxRetries > 5 {
		return nil, fmt.Errorf("maxRetries should be between 0 and 5, got %d", maxRetries)
	}

	// 验证 reTryTime 数组长度
	if len(reTryTime) != int(maxRetries) {
		return nil, fmt.Errorf("reTryTime array length must equal maxRetries, expected %d, got %d", maxRetries, len(reTryTime))
	}

	// 验证 reTryTime 中的负值
	for _, v := range reTryTime {
		if v < 0 {
			return nil, fmt.Errorf("reTryTime cannot contain negative values, got %v", reTryTime)
		}
	}

	// 排序 reTryTime 数组（如不需要可以删除这行）
	sort.Ints(reTryTime)

	// 返回初始化后的 Queue 实例
	return &BatchQueue{
		ctx:        context.Background(),
		processor:  processor,
		jobChan:    make(chan struct{}, 100),
		maxRetries: maxRetries,
		maxErr:     5,
		queue:      queue,
		failZSort:  queue + "_failed",
		log:        log.WithField("queue", queue),
		reTryTime:  reTryTime,
	}, nil
}

func (q *BatchQueue) Run() {
	defer q.recover()
	q.init()
	go q.reDo()
	for {
		select {
		case <-q.ctx.Done():
			return
		default:
			q.handleMaxError()
			msg, err := q.Pop()
			if err != nil {
				q.handleRedisError(err)
				continue
			}
			q.AddJob(msg)
			q.errCount = 0
		}
	}
}

func (q *BatchQueue) init() {
	if int(q.maxRetries) != len(q.reTryTime) {
		q.log.Error("队列初始化失败，重试配置不一致")
		return
	}

	q.reTryNow = q.isImmediateRetry()

	if q.ts < 1 {
		q.ts = 10
	}
}

func (q *BatchQueue) Pop() ([]string, error) {
	luaScript := `
		local result = {}
		for i = 1, ARGV[1] do
			local job = redis.call('RPOP', KEYS[1])
			if not job then
				break
			end
			table.insert(result, job)
		end
		return result
	`

	// 在 Redis 中执行 Lua 脚本
	result, err := utils.RedisClient.Eval(q.ctx, luaScript, []string{q.queue}, q.ts).Result()
	if err != nil {
		q.log.Errorf("Lua 脚本执行失败：%v", err)
		return nil, err
	}

	jobs, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	// 转换成字符串切片
	tasks := make([]string, len(jobs))
	for i, job := range jobs {
		tasks[i] = job.(string)
	}
	q.log.Infof("成功批量拉取 %d 个任务: %v", len(tasks), tasks)
	return tasks, nil
}

func (q *BatchQueue) Push(jobs []string) error {
	if err := utils.RedisClient.RPush(q.ctx, q.queue, jobs).Err(); err != nil {
		q.log.Error(err)
		return err
	}
	return nil
}

func (q *BatchQueue) AddJob(msg []string) {
	defer q.jobRecover(msg)
	q.jobChan <- struct{}{}
	go func() {
		defer q.JobDone()
		if err := q.processor.Run(q, msg); err != nil {
			q.log.Errorf("队列处理有误:%s，准备重新入队...", err)
			// 类型断言判断是否实现了 Retry 方法
			if retryProcessor, ok := q.processor.(interface{ Retry(*BatchQueue, []string) }); ok {
				retryProcessor.Retry(q, msg) // 调用 processor 自己的 Retry 方法
			} else {
				for _, v := range msg {
					q.Retry(v) // 调用通用的 Retry 方法
				}
			}
		}
	}()
}

func (q *BatchQueue) JobDone() {
	<-q.jobChan
}

func (q *BatchQueue) handleRedisError(err error) {
	if errors.Is(err, redis.Nil) {
		q.log.Info("队列暂无数据，等待中...")
		time.Sleep(5 * time.Second)
	} else {
		q.log.Errorf("Redis 错误: %v", err)
		q.errCount++
	}
}

func (q *BatchQueue) handleMaxError() {
	if q.errCount > q.maxErr {
		q.log.Warnf("连续错误超过最大次数，休眠 5 分钟")
		time.Sleep(5 * 60 * time.Second)
		q.errCount = 0 // 清零错误计数，避免死循环
	}
}

// Retry 自定义实现Retry(q *BatchQueue)，就调用自己的Retry
func (q *BatchQueue) Retry(msg string) {
	if msg == "" {
		return
	}

	var mpData = make(map[string]any)
	err := json.Unmarshal([]byte(msg), &mpData)
	if err != nil {
		q.log.Errorf("重试任务反序列化失败: %v", err)
		return
	}

	reTry, ok := mpData["re_try"].(float64)
	if !ok {
		reTry = 0
	}

	if int8(reTry) >= q.maxRetries {
		q.log.Warn("达到最大重试次数，不再重试")
		return
	}

	reTry++
	mpData["re_try"] = reTry
	data, err := json.Marshal(mpData)
	if err != nil {
		q.log.Errorf("重试任务序列化失败: %v", err)
		return
	}

	if q.reTryNow {
		err = utils.RedisClient.LPush(q.ctx, q.queue, data).Err()
	} else {
		err = utils.RedisClient.ZAdd(q.ctx, q.failZSort, &redis.Z{
			Score:  float64(time.Now().Add(q.calculateDelay(int(reTry))).Unix()),
			Member: data,
		}).Err()
	}

	if err != nil {
		q.log.Errorf("任务重新入队失败: %v", err)
	}
}

func (q *BatchQueue) isImmediateRetry() bool {
	for _, v := range q.reTryTime {
		if v > 0 {
			return false
		}
	}

	return true
}

func (q *BatchQueue) recover() {
	if err := recover(); err != nil {
		q.log.Errorf("Recovered panic in queue %s: %v  %s", q.queue, err, q.Data)
	}
}

func (q *BatchQueue) jobRecover(msg []string) {
	if err := recover(); err != nil {
		q.log.Errorf("jobRecover panic in queue %s: %v  %s", q.queue, err, q.Data)
		for _, v := range msg {
			q.Retry(v)
		}
	}
}

func (q *BatchQueue) calculateDelay(reTry int) time.Duration {
	if len(q.reTryTime) < reTry {
		return time.Minute
	}

	return time.Duration(q.reTryTime[reTry-1]) * time.Second
}

func (q *BatchQueue) reDo() {
	tickerInterval := time.Second * 10 // 默认的检查间隔时间
	ticker := time.NewTicker(tickerInterval)
	defer ticker.Stop()
	var lastTime int64
	pageSize := 30
	for range ticker.C {
		now := time.Now().Unix()
		minTime := fmt.Sprintf("%d", lastTime)
		maxTime := fmt.Sprintf("%d", now)
		fmt.Printf("正在检查%s ~ %s\n", minTime, maxTime)
		// 获取失败队列中的任务数量
		count, err := utils.RedisClient.ZCount(q.ctx, q.failZSort, minTime, maxTime).Result()
		if err != nil {
			q.log.Errorf("获取失败队列任务数量失败： %v", err)
			continue
		}

		page := int(math.Ceil(float64(count) / float64(pageSize)))
		for i := 0; i < page; i++ {
			Offset := i * pageSize
			result, err := utils.RedisClient.ZRangeByScore(q.ctx, q.failZSort, &redis.ZRangeBy{
				Min:    minTime,
				Max:    maxTime,
				Offset: int64(Offset),
				Count:  int64(pageSize),
			}).Result()

			if err != nil {
				q.log.Errorf("获取分数集合数据失败： %s", err)
				break
			}

			if err = q.Push(result); err != nil {
				q.log.Errorf("从新入队列失败： %s", err)
				break
			}
		}

		err = utils.RedisClient.ZRemRangeByScore(q.ctx, q.failZSort, minTime, maxTime).Err()
		if err != nil {
			q.log.Errorf("移除重新入队元素失败： %s", err)
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

func (q *BatchQueue) tickerInterval(count int64) int {
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

func (q *BatchQueue) Len() (int64, error) {
	return utils.RedisClient.LLen(q.ctx, q.queue).Result()
}
