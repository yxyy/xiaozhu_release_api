package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"strconv"
	"time"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/key"
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
)

type Loginer interface {
	verify() error
	login() (*user.MemberInfo, error)
}

type LoginLogic struct {
	ctx context.Context
	*Account
	*Mobile
	*WeChat
	*Email
	common.RequestForm
}

func NewLoginLogic(ctx context.Context) *LoginLogic {
	return &LoginLogic{
		ctx:     ctx,
		Account: &Account{ctx: ctx},
		Mobile:  &Mobile{ctx: ctx},
		WeChat:  &WeChat{ctx: ctx},
		Email:   &Email{ctx: ctx},
	}
}

type LoginResponse struct {
	UserId      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
	ExpireIn    int64  `json:"expire_in"`
	Username    string `json:"username"`
}

func NewLoginResponse() *LoginResponse {
	return &LoginResponse{}
}

// Login 登录控制
func (l *LoginLogic) Login(in Loginer) (resp *LoginResponse, err error) {
	if err = in.verify(); err != nil {
		return nil, err
	}

	// 执行对应的登录
	memberInfo, err := in.login()
	if err != nil {
		return nil, err
	}

	// 移除旧token信息
	err = l.RemoveToken(memberInfo.Id)
	if err != nil {
		return nil, err
	}

	// 校验账号情况
	if memberInfo.Status == 1 {
		return nil, errors.New("账号被封禁，禁止登录")
	}

	// 获取token信息
	resp, err = l.Token(memberInfo)
	if err != nil {
		return nil, err
	}

	// 登录信息入队列
	err = l.PushQueue()
	if err != nil {
		return nil, err
	}

	return

}

func (l *LoginLogic) Token(memberInfo *user.MemberInfo) (*LoginResponse, error) {

	response := NewLoginResponse()

	token, err := GetAccessToken(memberInfo)
	if err != nil {
		return nil, err
	}
	response.UserId = memberInfo.UserId
	response.AccessToken = token
	response.ExpireIn = int64(key.UserTokenExpress / 1000 / 1000 / 1000)

	// 缓存token信息
	keys := key.UserTokenPrefix + strconv.Itoa(memberInfo.UserId)
	err = utils.RedisClient.Set(l.ctx, keys, response.AccessToken, key.UserTokenExpress).Err()
	if err != nil {
		return nil, fmt.Errorf("token缓存设置失败：%s", err)
	}

	marshal, err := json.Marshal(&memberInfo)
	if err != nil {
		return nil, fmt.Errorf("序列化失败：%s", err)
	}

	err = utils.RedisClient.Set(l.ctx, response.AccessToken, marshal, key.UserTokenExpress).Err()
	if err != nil {
		return nil, fmt.Errorf("用户信息缓存设置失败：%s", err)
	}

	return response, nil
}

func (l *LoginLogic) RemoveToken(userId int) error {

	keys := key.UserTokenPrefix + strconv.Itoa(userId)
	token, err := utils.RedisClient.Get(l.ctx, keys).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}

		return fmt.Errorf("获取用户信息缓存失败：%s", err)
	}

	err = utils.RedisClient.Del(l.ctx, keys).Err()
	if err != nil {
		return fmt.Errorf("移除用户信息缓存token失败：%s", err)
	}

	err = utils.RedisClient.Del(l.ctx, token).Err()
	if err != nil {
		return fmt.Errorf("移除用户信息缓存失败：%s", err)
	}

	return nil

}

func (l *LoginLogic) PushQueue() error {
	l.RequestId = l.ctx.Value("request_id").(string)
	marshal, err := json.Marshal(&l)
	if err != nil {
		return fmt.Errorf("序列化登录信息失败：%s", err)
	}

	return utils.RedisClient.LPush(l.ctx, key.LoginQueue, marshal).Err()
}

func GetAccessToken(memberInfo *user.MemberInfo) (string, error) {

	claims := jwt.MapClaims{
		"user_id":  memberInfo.Id,
		"nickName": memberInfo.Nickname,
		"account":  memberInfo.Account,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(key.UserTokenExpress).Unix(),
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return Token.SignedString([]byte(viper.GetString("Auth.AccessSecret")))

}
