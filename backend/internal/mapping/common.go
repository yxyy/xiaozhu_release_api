package mapping

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
	"xiaozhu/utils"
)

func Get(key string) (map[int]string, error) {
	result, err := utils.RedisClient.Get(context.Background(), key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	if result != "" {
		data := make(map[int]string)
		if err = json.Unmarshal([]byte(result), &data); err != nil {
			return nil, err
		}

		return data, err
	}

	return nil, err
}

func Set(key string, data interface{}) error {

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return utils.RedisClient.Set(context.Background(), key, marshal, time.Second*360).Err()
}

func Remove(key string) error {

	return utils.RedisClient.Del(context.Background(), key).Err()
}
