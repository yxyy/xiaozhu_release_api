package handler

import (
	"xiaozhu/cron/internal/model/key"
	"xiaozhu/cron/internal/queue"
)

func StartQueue() {

	// 启动初始化队列
	go queue.NewQueue(key.InitQueue, queue.NewInitQueue()).Run()

	// 启动登录队列

}
