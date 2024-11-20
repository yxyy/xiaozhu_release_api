package mapping

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
	"xiaozhu/backend/internal/model/system"
	"xiaozhu/backend/utils"
)

func User() (userMap map[int]string, err error) {

	result, err := utils.RedisClient.Get(context.Background(), "userNameMap").Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return
	}
	userMap = make(map[int]string)
	if result != "" {
		if err = json.Unmarshal([]byte(result), &userMap); err != nil {
			return
		}

		return
	}

	user := system.NewSysUser()
	users, err := user.All()
	if err != nil {
		return
	}

	for _, v := range users {
		userMap[v.Id] = v.Nickname
	}
	marshal, err := json.Marshal(&userMap)
	if err != nil {
		return
	}

	err = utils.RedisClient.Set(context.Background(), "userNameMap", marshal, time.Second*60*60).Err()

	return
}
