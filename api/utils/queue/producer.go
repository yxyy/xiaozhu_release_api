package queue

import (
	"context"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

type Topic struct {
	Id      string `json:"id"`
	ReTry   int    `json:"re_try"`
	Message any    `json:"message"`
	Ts      int64  `json:"ts"`
}

func Uuid() string {
	return time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(8999)+1000)
}

func Push(ctx context.Context, key string, msg any) error {
	topic := Topic{
		Id:      Uuid(),
		ReTry:   0,
		Message: msg,
		Ts:      time.Now().Unix(),
	}

	marshal, err := json.Marshal(&topic)
	if err != nil {
		return err
	}

	return DefaultCoupler.Push(ctx, key, marshal)
}
