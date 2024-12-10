package job

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
	"xiaozhu/internal/model/key"
	"xiaozhu/internal/queue"
	"xiaozhu/utils"
	queue2 "xiaozhu/utils/queue"
)

const (
	maxQueue     = 3
	increase     = 3
	decrease     = -3
	maxThreshold = 1000
)

var (
	once           sync.Once
	DefaultMonitor *Monitor
)

func init() {
	once.Do(func() {
		DefaultMonitor = NewMonitor()
	})
}

type Monitor struct {
	ctx  context.Context
	list sync.Map
}

func NewMonitor() *Monitor {
	return &Monitor{
		ctx: context.Background(),
	}
}

func (m *Monitor) Run() {
	fmt.Println("监控开始")
	queueList := viper.GetStringMap("Queue")
	fmt.Println(queueList)
	if queueList == nil {
		fmt.Println("无队列配置")
		return
	}

	for k, v := range queueList {
		fmt.Println("正在查看队列:", k)
		queueName := QueueName(k)
		if queueName == "" {
			fmt.Println("无效队列队列配置")
			return
		}
		queueLen, err := m.QueueLen(queueName)
		if err != nil {
			fmt.Printf("获取队列长度失败%s\n", err)
			continue
		}
		fmt.Println("目前队列长度:", queueLen)

		queueInfo, ok := v.(map[string]any)
		if !ok {
			fmt.Println("解析队列配置失败")
			continue
		}

		go m.dispatch(queueName, queueLen, queueInfo)
	}

}

func (m *Monitor) QueueLen(queue string) (int, error) {
	result, err := utils.RedisDB00.LLen(m.ctx, queue).Result()
	return int(result), err
}

func (m *Monitor) getCancelFunc(k string) []context.CancelFunc {
	var cancelFunc []context.CancelFunc
	if val, ok := m.list.Load(k); ok {
		if tmp, ok := val.([]context.CancelFunc); ok {
			cancelFunc = tmp
		}
	}

	return cancelFunc
}

func (m *Monitor) getLast(k string) float64 {
	if val, ok := m.list.Load(k); ok {
		if tmp, ok := val.(float64); ok {
			return tmp
		}
	}

	return 0.0
}

func (m *Monitor) dispatch(queueName string, queueLen int, queueInfo map[string]any) {
	funcs := m.getCancelFunc(queueName)
	num := len(funcs)
	fmt.Println("当前活跃 goroutine 数量:", num)
	maxNum, threshold := getMaxQueueAndThresholdNum(queueInfo)
	thresholdLen := threshold * (num + 1)

	lastName := queueName + "_last"
	last := m.getLast(lastName)
	fmt.Println("上次last:", last)
	// 检查释放goroutine
	fmt.Println(queueLen, thresholdLen, num, maxNum, "-------------")
	if queueLen < thresholdLen && num > 0 {
		fmt.Println("进入减少goroutine")
		weight := float64(queueLen) / float64(thresholdLen)
		last = last - last*weight - 0.5
		fmt.Println("本次last:", last, weight, decrease)
		if last <= decrease {
			funcs[0]()
			m.list.Store(queueName, funcs[1:])
		}
		m.list.Store(lastName, last)
		return
	}

	// 检查释放需要新增goroutine
	if queueLen > thresholdLen && num < maxNum {
		fmt.Println("进入增加goroutine")
		weight := float64(queueLen) / float64(thresholdLen)
		last = last*weight + 0.5
		fmt.Println("本次last:", last, weight, increase)
		if last >= increase {
			fmt.Printf("队列:%s,长度:%d,超过阈值:%d了,加一个 goroutine,当前活跃 goroutine数量:%d\n", queueName, queueLen, thresholdLen, num)
			ctx, cancelFunc := context.WithCancel(m.ctx)
			q := NewMonitorQueue(ctx, queueName)
			fmt.Printf("%#v\n", q)
			if q == nil {
				cancelFunc()
				return
			}
			go q.Run()
			funcs = append(funcs, cancelFunc)
			m.list.Store(queueName, funcs)
		}
		m.list.Store(lastName, last)
		return
	}

	m.list.Store(lastName, 0)

}

func NewMonitorQueue(ctx context.Context, name string) *queue2.Queue {
	switch strings.ToUpper(name) {
	case strings.ToUpper(key.InitQueue):
		return queue2.NewQueueWithContext(ctx, key.InitQueue, queue.NewInit())
	default:
		return nil
	}
}

func getMaxQueueAndThresholdNum(queueInfo map[string]any) (int, int) {
	maxNum, ok := queueInfo["maxNum"].(int) // 运行开启runtime最大数量
	if !ok || maxNum <= 0 {
		maxNum = maxQueue
	}
	fmt.Println("最大运行 goruntine 数量:", maxNum)

	threshold, ok := queueInfo["threshold"].(int)
	if !ok || threshold <= 0 {
		threshold = maxThreshold
	}
	fmt.Println("队列阈值是:", threshold)

	return maxNum, threshold
}

func QueueName(name string) string {
	switch strings.ToUpper(name) {
	case strings.ToUpper(key.InitQueue):
		return key.InitQueue
	default:
		return ""
	}
}
