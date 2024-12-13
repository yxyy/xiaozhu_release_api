package handler

import (
	"fmt"
	"github.com/spf13/viper"
	"xiaozhu/internal/queue"
	"xiaozhu/utils"
	utilsqueue "xiaozhu/utils/queue"
)

func StartQueue() {
	InitQueue()
	// 启动初始化队列
	go queue.NewInitQueue().Run()

	go queue.NewLoginQueue().Run()
	// 启动登录队列

}

func InitQueue() {
	// 注册队列连接器
	redis := &utilsqueue.Redis{Conn: utils.RedisDB00}
	utilsqueue.RegisterCoupler(redis)

	// 注册监控配置
	queueInfo := viper.GetStringMap("Queue")
	var config []*utilsqueue.Config
	for k, v := range queueInfo {
		qv, ok := v.(map[string]any)
		if !ok {
			fmt.Println("解析队列配置失败")
			continue
		}
		maxNum, ok := qv["maxnum"].(int) // 运行开启runtime最大数量
		if !ok || maxNum <= 0 {
			continue
		}

		threshold, ok := qv["threshold"].(int) // 运行开启runtime最大数量
		if !ok || threshold <= 0 {
			continue
		}

		config = append(config, &utilsqueue.Config{
			Name:        k,
			MaxQueueNum: maxNum,
			Threshold:   threshold,
		})
	}

	utilsqueue.RegisterMonitorConfig(config)

	utilsqueue.StartMonitor()

}
