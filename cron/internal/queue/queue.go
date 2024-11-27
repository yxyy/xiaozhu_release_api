package queue

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
	"xiaozhu/cron/utils"
)

type Processor interface {
	Run(*Queue, string) error
}

type Queue struct {
	ctx        context.Context
	processor  Processor
	jobChan    chan struct{}
	maxRetries int8
	log        *log.Entry
	errCount   int
	maxErr     int
	queue      string
	Data       string
}

func NewQueue(queue string, processor Processor) *Queue {
	return &Queue{
		ctx:        context.Background(),
		processor:  processor,
		jobChan:    make(chan struct{}, 100),
		maxRetries: 3,
		maxErr:     5,
		queue:      queue,
		log:        log.WithField("queue", queue),
	}
}

func (q *Queue) Run() {
	defer q.recover()
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

func (q *Queue) Pop() (string, error) {
	result, err := utils.RedisClient.BRPop(q.ctx, time.Second*5, q.queue).Result()
	if err != nil {
		return "", err
	}
	q.Data = result[1]
	q.log.Info(result[1])

	return result[1], err
}

func (q *Queue) AddJob(msg string) {
	q.jobChan <- struct{}{}
	go func() {
		defer q.JobDone()
		if err := q.processor.Run(q, msg); err != nil {
			q.log.Errorf("队列处理有误:%s，准备重新入队...", err)
			// 类型断言判断是否实现了 Retry 方法
			if retryProcessor, ok := q.processor.(interface{ Retry(*Queue) }); ok {
				retryProcessor.Retry(q) // 调用 processor 自己的 Retry 方法
			} else {
				q.Retry() // 调用通用的 Retry 方法
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
func (q *Queue) Retry() {
	if q.Data == "" {
		return
	}

	var mpData = make(map[string]any)
	err := json.Unmarshal([]byte(q.Data), &mpData)
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

	if err = utils.RedisClient.LPush(q.ctx, q.queue, data).Err(); err != nil {
		q.log.Errorf("任务重新入队失败: %v", err)
	}
}

func (q *Queue) recover() {
	if err := recover(); err != nil {
		q.log.Errorf("Recovered panic in queue %s: %v  %s", q.queue, err, q.Data)
	}
}
