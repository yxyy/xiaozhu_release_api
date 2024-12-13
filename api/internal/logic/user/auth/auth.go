package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"xiaozhu/internal/config/cache"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/key"
	"xiaozhu/internal/model/user"
	"xiaozhu/utils"
	"xiaozhu/utils/queue"
)

type Auther interface {
	verify() error
	login() (*user.MemberInfo, error)
	register(common.RequestForm) (*user.MemberInfo, error)
}

type Logic struct {
	ctx context.Context
	*Account
	*Mobile
	*WeChat
	*Email
	common.RequestForm
}

func NewAuthLogic(ctx context.Context) *Logic {
	return &Logic{
		ctx:     ctx,
		Account: NewAccount(ctx),
		Mobile:  NewMobile(ctx),
		WeChat:  NewWeChat(ctx),
		Email:   NewEmail(ctx),
	}
}

func NewAuther(l *Logic) (Auther, error) {
	switch {
	case l.Mobile.Phone != 0: //  手机登录
		return l.Mobile, nil
	case l.Email.Email != "": // 邮箱登录
		return l.Email, nil
	case l.WeChat.WxCode != "": // 微信登录
		return l.WeChat, nil
	case l.Account.Account != "": // 账号登录
		return l.Account, nil
	default:
		return nil, errors.New("无效的登录方式")
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
func (l *Logic) Login(in Auther) (resp *LoginResponse, err error) {
	if err = in.verify(); err != nil {
		return nil, err
	}

	// 执行对应的登录
	memberInfo, err := in.login()
	if err != nil {
		return nil, err
	}

	// 移除旧token信息
	if err = l.RemoveToken(memberInfo.Id); err != nil {
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
	if err = l.PushLoginQueue(); err != nil {
		return nil, err
	}

	return

}

func (l *Logic) Register(in Auther) (resp *user.MemberInfo, err error) {
	if err = in.verify(); err != nil {
		return nil, err
	}

	// 执行对应的登录
	memberInfo, err := in.register(l.RequestForm)
	if err != nil {
		return nil, err
	}

	// 登录信息入队列
	err = l.PushRegisterQueue(memberInfo.Id)
	if err != nil {
		return nil, err
	}

	return memberInfo, nil

}

func (l *Logic) Token(memberInfo *user.MemberInfo) (*LoginResponse, error) {

	response := NewLoginResponse()
	response.UserId = memberInfo.UserId
	response.AccessToken = GetAccessToken(memberInfo)
	response.ExpireIn = int64(key.UserTokenExpress / 1000 / 1000 / 1000)

	// 缓存token信息
	keys := key.UserTokenPrefix + strconv.Itoa(memberInfo.UserId)
	err := cache.RedisDB00.Set(l.ctx, keys, response.AccessToken, key.UserTokenExpress).Err()
	if err != nil {
		return nil, fmt.Errorf("token缓存设置失败：%s", err)
	}

	marshal, err := json.Marshal(&memberInfo)
	if err != nil {
		return nil, fmt.Errorf("序列化失败：%s", err)
	}

	err = cache.RedisDB00.Set(l.ctx, key.LoginTokenPrefix+response.AccessToken, marshal, key.UserTokenExpress).Err()
	if err != nil {
		return nil, fmt.Errorf("用户信息缓存设置失败：%s", err)
	}

	return response, nil
}

func (l *Logic) RemoveToken(userId int) error {
	keys := key.UserTokenPrefix + strconv.Itoa(userId)
	fmt.Println(keys)
	token, err := cache.RedisDB00.Get(l.ctx, keys).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return fmt.Errorf("获取用户信息缓存失败：%s", err)
	}
	fmt.Println(token)
	if err = cache.RedisDB00.Del(l.ctx, keys).Err(); err != nil {
		return fmt.Errorf("移除用户信息缓存token失败：%s", err)
	}
	fmt.Println(key.LoginTokenPrefix + token)
	if err = cache.RedisDB00.Del(l.ctx, key.LoginTokenPrefix+token).Err(); err != nil {
		fmt.Println(err)
		return fmt.Errorf("移除用户信息缓存失败：%s", err)
	}

	return nil

}

func (l *Logic) PushLoginQueue() error {
	l.RequestId = l.ctx.Value("request_id").(string)
	// marshal, err := json.Marshal(&l)
	// if err != nil {
	// 	return fmt.Errorf("序列化登录信息失败：%s", err)
	// }

	return queue.Push(l.ctx, key.LoginQueue, l)

}

func (l *Logic) PushRegisterQueue(userId int) error {
	l.RequestId = l.ctx.Value("request_id").(string)
	l.UserId = userId
	// marshal, err := json.Marshal(&l)
	// if err != nil {
	// 	return fmt.Errorf("序列化登录信息失败：%s", err)
	// }

	return cache.RedisDB00.LPush(l.ctx, key.RegisterQueue, l).Err()
}

func GetAccessToken(memberInfo *user.MemberInfo) string {

	return utils.Md5(memberInfo.Account + utils.Random(10))

}
