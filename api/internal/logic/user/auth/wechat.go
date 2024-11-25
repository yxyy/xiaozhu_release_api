package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"time"
	"xiaozhu/api/internal/logic/common"
	"xiaozhu/api/internal/model/user"
	"xiaozhu/api/utils"
)

type WeChat struct {
	ctx    context.Context
	WxCode string `json:"wx_code" form:"wx_code" gorm:"wx_code"`
}

func NewWeChat(ctx context.Context) *WeChat {
	return &WeChat{ctx: ctx}
}

func (w *WeChat) verify() error {
	if w.WxCode == "" {
		return errors.New("登录码不能为空")
	}
	return nil
}

func (w *WeChat) login() (memberInfo *user.MemberInfo, err error) {

	params := url.Values{}
	params.Add("appid", viper.GetString("mini.appid"))
	params.Add("secret", viper.GetString("mini.secret"))
	params.Add("js_code", w.WxCode)
	params.Add("grant_type", "authorization_code")

	urls := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?%s", params.Encode())

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(urls)
	if err != nil {
		return nil, fmt.Errorf("请求微信API失败: %w", err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	var response WxLoginResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, fmt.Errorf("解析微信响应失败: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("微信登录失败: %s (错误代码: %d)", response.ErrMsg, response.ErrCode)
	}

	return response.findOrCreateUserByOpenid()
}

type WxLoginResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	Openid     string `json:"openid"`
}

func (w *WxLoginResponse) findOrCreateUserByOpenid() (memberInfo *user.MemberInfo, err error) {
	err = utils.MysqlDb.Model(&memberInfo).Where("wechat", w.Openid).First(&memberInfo).Error
	if err == nil {
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	unix := time.Now().Unix()
	memberInfo.Account = "wx_" + utils.Random(8)
	memberInfo.Salt = utils.Salt()
	memberInfo.Password = utils.Md5SaltAndPassword(memberInfo.Salt, memberInfo.Salt)
	memberInfo.Wechat = w.Openid
	memberInfo.LastLoginTime = unix
	memberInfo.UpdatedAt = unix
	memberInfo.CreatedAt = unix

	return memberInfo, utils.MysqlDb.Model(&memberInfo).Create(&memberInfo).Error
}

func (w *WeChat) register(req common.RequestForm) (memberInfo *user.MemberInfo, err error) {

	params := url.Values{}
	params.Add("appid", viper.GetString("mini.appid"))
	params.Add("secret", viper.GetString("mini.secret"))
	params.Add("js_code", w.WxCode)
	params.Add("grant_type", "authorization_code")

	urls := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?%s", params.Encode())

	client := http.Client{Timeout: time.Second * 10}
	resp, err := client.Get(urls)
	if err != nil {
		return nil, fmt.Errorf("请求微信API失败: %w", err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	var response WxLoginResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, fmt.Errorf("解析微信响应失败: %w", err)
	}
	if response.ErrCode != 0 {
		return nil, fmt.Errorf("微信登录失败: %s (错误代码: %d)", response.ErrMsg, response.ErrCode)
	}

	return response.findOrCreateUserByOpenid()
}
