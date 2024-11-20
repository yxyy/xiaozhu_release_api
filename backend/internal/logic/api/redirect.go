package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"time"
	"xiaozhu/backend/internal/model/market/assets"
	"xiaozhu/backend/utils"
)

type BmOAuthRedirectRequest struct {
	State    string `json:"state" form:"state"`
	AuthCode string `json:"auth_code" form:"auth_code"`
}

type BmLogic struct {
	ctx context.Context
	BmOAuthRedirectRequest
	MarketAppId assets.MarketAppId
}

func NewBmLogic(ctx context.Context) *BmLogic {
	return &BmLogic{ctx: ctx}
}

func (l *BmLogic) Do() error {
	if l.State == "" {
		return errors.New("state 参数缺失")
	}
	if l.AuthCode == "" {
		return errors.New("auth_code 参数缺失")
	}

	l.MarketAppId.State = l.State
	if err := l.MarketAppId.Get(l.ctx); err != nil {
		return err
	}

	if l.MarketAppId.AppId == "" || l.MarketAppId.Secret == "" {
		return errors.New("AppId 或者 Secret 缺失")
	}

	req := &AccessTokenRequest{
		AppId:    l.MarketAppId.AppId,
		Secret:   l.MarketAppId.Secret,
		AuthCode: l.AuthCode,
	}

	token, err := l.AccessToken(req)
	if err != nil {
		return err
	}

	fmt.Println("-----------水水水水水水水水水水水------------", token, err)

	if token.Code != 0 {
		return errors.New("授权失败：" + token.Message)
	}

	token.Data.LastTime = time.Now().Unix()
	token.Data.ExpiresIn = token.Data.LastTime + token.Data.ExpiresIn

	marshal, err := json.Marshal(&token.Data)
	if err != nil {
		return err
	}

	key := viper.GetString("bm_access_token")
	if err = utils.RedisClient.HSet(l.ctx, key, l.State, string(marshal)).Err(); err != nil {
		return err
	}

	return nil
}

type AccessTokenRequest struct {
	AppId    string `json:"app_id"`
	Secret   string `json:"secret"`
	AuthCode string `json:"auth_code"`
}

type Data struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AdvertiserId          int    `json:"advertiser_id"`
	AdvertiserIds         []int  `json:"advertiser_ids"`
	ExpiresIn             int64  `json:"expires_in"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	LastTime              int64  `json:"last_time"`
}

type AccessTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func (l *BmLogic) AccessToken(in *AccessTokenRequest) (accessToken *AccessTokenResponse, err error) {
	if in == nil {
		return nil, errors.New("无效参数")
	}

	bodyByte, err := json.Marshal(&in)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bodyByte))

	body := bytes.NewReader(bodyByte)
	url := viper.GetString("Bm.Host") + viper.GetString("Bm.AccessToken")
	fmt.Println(url)

	// response, err := http.Post(url, "application/json", body)

	response, err := utils.Request(l.ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------------------------", response.StatusCode, response)
	defer func(Body io.ReadCloser) {
		if err2 := Body.Close(); err2 != nil {
			log.Error(err2)
		}
	}(response.Body)

	// accessToken = new(AccessTokenResponse)
	// decoder := json.NewDecoder(response.Body)
	// err = decoder.Decode(&accessToken)
	// if err != nil {
	// 	return nil, err
	// }

	responseBodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("请求失败，状态码:%d，响应内容：%s", response.StatusCode, responseBodyByte))
	}

	err = json.Unmarshal(responseBodyByte, &accessToken)
	if err != nil {
		return nil, err
	}

	return accessToken, nil

}
