package queue

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

const (
	increase     = 3
	decrease     = -3
	maxQueueNum  = 2
	maxThreshold = 1000
)

var (
	once     sync.Once
	QMonitor *Monitor
	QConfig  = make(map[string]*Config)
)

func init() {
	once.Do(func() {
		QMonitor = NewMonitor()
	})
}

type Monitor struct {
	ctx   context.Context
	state sync.Map          // 运行状态，取消函数，步数
	List  map[string]*Queue // 注册的监控队列
	mu    sync.Mutex
}

type Config struct {
	Name        string
	MaxQueueNum int
	Threshold   int
}

func StartMonitor() {
	go QMonitor.Run()
}

// RegisterMonitor 注入监控队列
func RegisterMonitor(name string, q *Queue) {
	if name == "" || q == nil {
		return
	}

	fmt.Println("注册队列：", name)

	name = strings.ToUpper(name)
	QMonitor.mu.Lock()
	defer QMonitor.mu.Unlock()

	if QMonitor.List == nil {
		QMonitor.List = make(map[string]*Queue)
	}
	if QMonitor.List[name] == nil {
		QMonitor.List[name] = q
	}
	if QConfig[name] == nil {
		QConfig[name] = &Config{
			Name:        name,
			MaxQueueNum: maxQueueNum,
			Threshold:   maxThreshold,
		}
	}
}

// RegisterMonitorConfig 注入监控配置
func RegisterMonitorConfig(config []*Config) {
	if len(config) == 0 {
		return
	}

	for _, v := range config {
		name := strings.ToUpper(v.Name)
		if name == "" {
			continue
		}
		if v.MaxQueueNum <= 0 {
			v.MaxQueueNum = maxQueueNum
		}
		if v.Threshold <= 0 {
			v.Threshold = maxThreshold
		}
		QConfig[name] = &Config{
			Name:        name,
			MaxQueueNum: v.MaxQueueNum,
			Threshold:   v.Threshold,
		}
	}
}

func NewMonitor() *Monitor {
	return &Monitor{
		ctx:  context.Background(),
		List: make(map[string]*Queue),
	}
}

func (m *Monitor) Run() {
	fmt.Println("监控队列开始")
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for name, q := range QMonitor.List {
			queueLen, err := q.Len()
			if err != nil {
				fmt.Printf("获取队列:%s,长度失败%s\n", name, err)
				continue
			}
			fmt.Printf("队列:%s,长度:%d\n", name, queueLen)

			config := *QConfig[name]
			if config.Name == "" {
				config = Config{
					Name:        name,
					MaxQueueNum: maxQueueNum,
					Threshold:   maxThreshold,
				}
			}

			go m.dispatch(name, queueLen, config, q)
		}
	}
}

func (m *Monitor) getCancelFunc(k string) []context.CancelFunc {
	var cancelFunc []context.CancelFunc
	if val, ok := m.state.Load(k); ok {
		if tmp, ok := val.([]context.CancelFunc); ok {
			cancelFunc = tmp
		}
	}

	return cancelFunc
}

func (m *Monitor) getLast(k string) float64 {
	if val, ok := m.state.Load(k); ok {
		if tmp, ok := val.(float64); ok {
			return tmp
		}
	}

	return 0.0
}

func (m *Monitor) dispatch(name string, queueLen int64, config Config, q *Queue) {
	funcs := m.getCancelFunc(name)
	num := len(funcs)
	// maxNum, threshold := getMaxQueueAndThresholdNum(queueInfo)
	thresholdLen := int64(config.Threshold * (num + 1))
	fmt.Printf("当前活跃 goroutine 数量:%d,队列长度是：%d,阈值是：%d \n", num, queueLen, thresholdLen)
	logrus.WithField("action", "monitor").Infof("当前活跃 goroutine 数量:%d,队列长度是：%d,阈值是：%d \n", num, queueLen, thresholdLen)

	lastName := name + "_last"
	last := m.getLast(lastName)
	fmt.Println("上次last:", last)
	// 检查释放goroutine
	if queueLen < thresholdLen && num > 0 {
		fmt.Println("进入减少goroutine")
		weight := float64(queueLen) / float64(thresholdLen)
		last = last - last*weight - 0.5
		fmt.Println("本次last:", last, weight, decrease)
		if last <= decrease {
			funcs[0]()
			m.state.Store(name, funcs[1:])
			last = 0
		}
		m.state.Store(lastName, last)
		return
	}

	// 检查释放需要新增goroutine
	if queueLen > thresholdLen && num < config.MaxQueueNum {
		fmt.Println("进入增加goroutine")
		weight := float64(queueLen) / float64(thresholdLen)
		last = last*weight + 0.5
		fmt.Println("本次last:", last, weight, increase)
		if last >= increase {
			fmt.Printf("队列:%s,长度:%d,超过阈值:%d了,加一个 goroutine,当前活跃 goroutine数量:%d\n", name, queueLen, thresholdLen, num)
			ctx, cancelFunc := context.WithCancel(m.ctx)
			nq := CopyQueue(ctx, q)
			// fmt.Printf("%#v\n", nq)
			if q == nil {
				cancelFunc()
				return
			}
			go nq.Run()
			funcs = append(funcs, cancelFunc)
			m.state.Store(name, funcs)
			last = 0
		}
		m.state.Store(lastName, last)
		return
	}

	m.state.Store(lastName, 0)

}

func CopyQueue(ctx context.Context, old *Queue) *Queue {
	return &Queue{
		Ctx:            ctx,
		processor:      old.processor,
		batchProcessor: old.batchProcessor,
		Coupler:        old.Coupler,
		jobChan:        make(chan struct{}, jobChanCount),
		Log:            old.Log,
		errCount:       old.errCount,
		maxErr:         old.maxErr,
		name:           old.name,
		failName:       old.failName,
		maxRetries:     old.maxRetries,
		reTryTime:      old.reTryTime,
		reTryNow:       old.reTryNow,
		ts:             old.ts,
	}
}
