package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	RedisDB00 *redis.Client
	// RedisDB01 *redis.Client
	// RedisDB02 *redis.Client
)

func InitRedis() error {
	RedisDB00 = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.master.host") + ":" + viper.GetString("redis.master.port"),
		Password: viper.GetString("redis.master.password"),
		DB:       viper.GetInt("redis.master.db"),
	})
	_, err := RedisDB00.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	// RedisDB01 = redis.NewClient(&redis.Options{
	// 	Addr:     viper.GetString("redis.master.host") + ":" + viper.GetString("redis.master.port"),
	// 	Password: viper.GetString("redis.master.password"),
	// 	DB:       viper.GetInt("redis.master.db"),
	// })
	// _, err = RedisDB01.Ping(context.Background()).Result()
	// if err != nil {
	// 	return err
	// }
	//
	// RedisDB02 = redis.NewClient(&redis.Options{
	// 	Addr:     viper.GetString("redis.master.host") + ":" + viper.GetString("redis.master.port"),
	// 	Password: viper.GetString("redis.master.password"),
	// 	DB:       viper.GetInt("redis.master.db"),
	// })
	// _, err = RedisDB02.Ping(context.Background()).Result()
	// if err != nil {
	// 	return err
	// }

	return nil
}
