package user

import (
	"context"
	"errors"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/utils/queue"
)

type CaptchaLogic struct {
	ctx context.Context
	CaptchaRequest
}

func NewCaptchaLogic(ctx context.Context) *CaptchaLogic {
	return &CaptchaLogic{ctx: ctx}
}

type CaptchaRequest struct {
	common.RequestForm
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (l *CaptchaLogic) Captcha() error {
	if l.CaptchaRequest.Email == "" && l.CaptchaRequest.Phone == "" {
		return errors.New("邮箱或者电话都为空")
	}
	return queue.Push(l.ctx, key.CodeQueue, l.CaptchaRequest)
}
