package user

import (
	"context"
	"encoding/json"
	"errors"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/assets"
	"xiaozhu/api/internal/model/key"
	"xiaozhu/api/utils"
)

type InitLogic struct {
	ctx         context.Context
	InitRequest InitRequest
}

type InitRequest struct {
	common.RequestForm
	ProductKey string `json:"product_key" binding:"required"`
	Lang       string `json:"lang"`
	Debug      string `json:"debug" `
}

type InitResponse struct {
	Protocol        string `json:"protocol"`          // 协议
	Privacy         string `json:"privacy"`           // 隐私
	OpenMailCode    int8   `json:"open_mail_code"`    // 邮箱验证
	OpenMobileCode  int8   `json:"open_mobile_code"`  // 手机验证
	UserCenterUrl   string `json:"user_center_url"`   // 用户中心连接
	UserCenterImage string `json:"user_center_image"` // 用户中心悬浮图标
	IsAuthRealName  int8   `json:"is_auth_real_name"`
	IsLimitUnderage int8   `json:"is_limit_underage"`
}

func NewInitLogic(ctx context.Context) *InitLogic {
	return &InitLogic{ctx: ctx}
}

func (l *InitLogic) Init() (*InitResponse, error) {

	appGame, err := assets.GetAppGameInfo(l.ctx, l.InitRequest.GameId)
	if err != nil {
		return nil, err
	}

	if appGame.Status > 1 {
		return nil, errors.New("游戏已经下架")
	}

	// todo 日志入队列
	l.InitRequest.RequestId = l.ctx.Value("request_id").(string)

	marshal, err := json.Marshal(&l.InitRequest)
	if err != nil {
		return nil, err
	}

	if err = utils.RedisClient.LPush(l.ctx, key.InitQueue, marshal).Err(); err != nil {
		return nil, err
	}

	var data = &InitResponse{}
	data.Privacy = appGame.GameName
	data.IsAuthRealName = appGame.IsAuthRealName
	data.IsLimitUnderage = appGame.IsLimitUnderage

	return data, nil
}

func (l *InitLogic) Verify() error {
	if l.InitRequest.GameId == 0 {
		return errors.New("游戏不能为空")
	}

	if l.InitRequest.CpCode == "" {
		return errors.New("cp标识不能为空")
	}

	if l.InitRequest.Os == "" {
		return errors.New("操作系统不能为空")
	}

	if l.InitRequest.Version == "" {
		return errors.New("版本不能为空")
	}

	if l.InitRequest.DeviceId == "" {
		return errors.New("设备不能为空")
	}

	if l.InitRequest.ProductKey == "" {
		return errors.New("产品密钥不能为空")
	}

	return nil
}
