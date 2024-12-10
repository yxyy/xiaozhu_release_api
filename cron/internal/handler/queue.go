package handler

import (
	"xiaozhu/internal/queue"
)

func StartQueue() {

	// 启动初始化队列
	go queue.NewInitQueue().Run()

	// 启动登录队列

}
