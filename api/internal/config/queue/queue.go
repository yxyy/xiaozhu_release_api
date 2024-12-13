package queue

import (
	"xiaozhu/internal/config/cache"
	utilsQ "xiaozhu/utils/queue"
)

func Init() error {

	if cache.RedisDB00 == nil {
		if err := cache.InitRedis(); err != nil {
			return err
		}
	}

	// 注册队列连接器
	redis := &utilsQ.Redis{Conn: cache.RedisDB00}
	utilsQ.RegisterCoupler(redis)

	return nil
}
