package queue

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Coupler interface {
	Len(ctx context.Context, key string) (int64, error)
	Pop(ctx context.Context, key string) ([]string, error)
	BatchPop(ctx context.Context, key string, ts int) ([]string, error)
	Push(ctx context.Context, key string, msg any) error
	FailAdd(ctx context.Context, key string, expiredScore float64, msg any) error
	FailNum(ctx context.Context, key, prev, next string) (int64, error)
	FailRangeByScore(ctx context.Context, key, minTime, maxTime string, offset, count int64) ([]string, error)
	FailRemRangeByScore(ctx context.Context, key, minTime, maxTime string) error
}

var DefaultCoupler Coupler

func RegisterCoupler(coupler Coupler) {
	DefaultCoupler = coupler
}

type Redis struct {
	Conn *redis.Client
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisCouplerByConfig(config RedisConfig) (Coupler, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{Conn: redisClient}, nil
}

func (r *Redis) Len(ctx context.Context, key string) (int64, error) {
	return r.Conn.LLen(ctx, key).Result()
}

func (r *Redis) Pop(ctx context.Context, key string) ([]string, error) {
	result, err := r.Conn.BRPop(ctx, time.Second*5, key).Result()
	if err != nil {
		return nil, err
	}

	return []string{result[1]}, err
}

func (r *Redis) BatchPop(ctx context.Context, key string, ts int) ([]string, error) {
	luaScript := `
		local result = {}
		for i = 1, ARGV[1] do
			local job = redis.call('RPOP', KEYS[1])
			if not job then
				break
			end
			table.insert(result, job)
		end
		return result
	`

	// 在 Redis 中执行 Lua 脚本
	result, err := r.Conn.Eval(ctx, luaScript, []string{key}, ts).Result()
	if err != nil {
		return nil, err
	}

	jobs, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	// 转换成字符串切片
	tasks := make([]string, len(jobs))
	for i, job := range jobs {
		tasks[i] = job.(string)
	}
	return tasks, nil
}

func (r *Redis) Push(ctx context.Context, key string, jobs any) error {
	if err := r.Conn.LPush(ctx, key, jobs).Err(); err != nil {
		return err
	}
	return nil
}

func (r *Redis) FailAdd(ctx context.Context, key string, expiredScore float64, msg any) error {
	return r.Conn.ZAdd(ctx, key, &redis.Z{
		Score:  expiredScore,
		Member: msg,
	}).Err()
}

func (r *Redis) FailNum(ctx context.Context, key, minTime, maxTime string) (int64, error) {
	return r.Conn.ZCount(ctx, key, minTime, maxTime).Result()
}

func (r *Redis) FailRangeByScore(ctx context.Context, key, minTime, maxTime string, offset, count int64) ([]string, error) {
	return r.Conn.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min:    minTime,
		Max:    maxTime,
		Offset: offset,
		Count:  count,
	}).Result()
}

func (r *Redis) FailRemRangeByScore(ctx context.Context, key, minTime, maxTime string) error {
	return r.Conn.ZRemRangeByScore(ctx, key, minTime, maxTime).Err()
}
