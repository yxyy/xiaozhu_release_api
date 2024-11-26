package handler

import (
	"context"
	"xiaozhu/cron/internal/queue"
)

func StartQueue() {

	// 启动初始化队列
	go queue.Init(context.Background())

	// 启动登录队列

}
