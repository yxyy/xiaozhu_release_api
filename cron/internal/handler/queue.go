package handler

import (
	"xiaozhu/internal/model/key"
	"xiaozhu/internal/queue"
)

func StartQueue() {

	// 启动初始化队列
	go queue.NewQueue(key.InitQueue, queue.NewInitQueue()).Run()
	// go queue.NewQueue(key.InitQueue, queue.NewInitQueue()).Run()
	// go queue.NewQueue(key.InitQueue, queue.NewInitQueue()).Run()
	// 启动登录队列

}
