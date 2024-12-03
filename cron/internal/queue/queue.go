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

type Processor interface {
	Run(*Queue, string) error
}

type Queue struct {
	ctx        context.Context
	processor  Processor
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
}

func NewQueue(queue string, processor Processor) *Queue {
	return &Queue{
		ctx:        context.Background(),
		processor:  processor,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		queue:      queue,
		failZSort:  queue + "_failed",
		log:        log.WithField("queue", queue),
		reTryTime:  []int{60, 300, 1800},
	}
}

func NewRetryQueue(queue string, processor Processor, maxRetries int8, reTryTime []int) (*Queue, error) {
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
	return &Queue{
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

func (q *Queue) Run() {
	defer q.recover()
	q.init()
	go q.ReDo()
	for {
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

func (q *Queue) init() {
	if int(q.maxRetries) != len(q.reTryTime) {
		q.log.Error("队列初始化失败，重试配置不一致")
		return
	}

	q.reTryNow = q.isImmediateRetry()
}

func (q *Queue) Pop() (string, error) {
	result, err := utils.RedisClient.BRPop(q.ctx, time.Second*5, q.queue).Result()
	if err != nil {
		return "", err
	}
	q.Data = result[1]
	q.log.Info(result[1])

	return result[1], err
}

func (q *Queue) Push(jobs []string) error {
	if err := utils.RedisClient.RPush(q.ctx, q.queue, jobs).Err(); err != nil {
		q.log.Error(err)
		return err
	}
	return nil
}

func (q *Queue) AddJob(msg string) {
	defer q.jobRecover(msg)
	q.jobChan <- struct{}{}
	go func() {
		defer q.JobDone()
		if err := q.processor.Run(q, msg); err != nil {
			q.log.Errorf("队列处理有误:%s，准备重新入队...", err)
			// 类型断言判断是否实现了 Retry 方法
			if retryProcessor, ok := q.processor.(interface{ Retry(*Queue, string) }); ok {
				retryProcessor.Retry(q, msg) // 调用 processor 自己的 Retry 方法
			} else {
				q.Retry(msg) // 调用通用的 Retry 方法
			}
		}
	}()
}

func (q *Queue) JobDone() {
	<-q.jobChan
}

func (q *Queue) handleRedisError(err error) {
	if errors.Is(err, redis.Nil) {
		q.log.Info("队列暂无数据，等待中...")
		time.Sleep(5 * time.Second)
	} else {
		q.log.Errorf("Redis 错误: %v", err)
		q.errCount++
	}
}

func (q *Queue) handleMaxError() {
	if q.errCount > q.maxErr {
		q.log.Warnf("连续错误超过最大次数，休眠 5 分钟")
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
		q.log.Errorf("Recovered panic in queue %s: %v  %s", q.queue, err, q.Data)
	}
}

func (q *Queue) jobRecover(msg string) {
	if err := recover(); err != nil {
		q.log.Errorf("jobRecover panic in queue %s: %v  %s", q.queue, err, q.Data)
		q.Retry(msg)
	}
}

func (q *Queue) calculateDelay(reTry int) time.Duration {
	if len(q.reTryTime) < reTry {
		return time.Minute
	}

	return time.Duration(q.reTryTime[reTry-1]) * time.Second
}

func (q *Queue) ReDo() {
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
