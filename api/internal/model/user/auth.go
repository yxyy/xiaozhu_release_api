package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

type LoginResponse struct {
	AccessToken   string   `json:"access_token"`
	RefreshToken  string   `json:"refresh_token"`
	AccessExpire  int64    `json:"access_expire"`
	RefreshExpire int64    `json:"refresh_expire"`
	Role          []string `json:"role"`
	Username      string   `json:"username"`
}

func NewLoginResponse() *LoginResponse {
	return &LoginResponse{}
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" gorm:"refresh_token" form:"refresh_token"`
}

type RoleMenu struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parentId"`
	MenuType int8   `json:"menuType"`
	Title    string `json:"title"`
}

func NewRefreshRequest() *RefreshRequest {
	return &RefreshRequest{}
}

func GetAccessToken(user *SysUser) (string, error) {

	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"nickName": user.Nickname,
		"role":     user.RoleIds,
		"account":  user.Account,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * time.Duration(viper.GetInt64("Auth.AccessExpire"))).Unix(),
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return Token.SignedString([]byte(viper.GetString("Auth.AccessSecret")))

}

func GetRefreshToken(user *SysUser) (string, error) {

	claims := jwt.MapClaims{
		"user_id":  user.Id,
		"nickName": user.Nickname,
		"role":     user.RoleIds,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * time.Duration(viper.GetInt64("Auth.RefreshExpire"))).Unix(),
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return Token.SignedString([]byte(viper.GetString("Auth.RefreshSecret")))

}

func ParseToken(token string, tokenType int) (*SysUser, error) {
	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		secret := ""
		if tokenType == 1 {
			secret = viper.GetString("Auth.AccessSecret")
		} else {
			secret = viper.GetString("Auth.RefreshSecret")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !parse.Valid {
		return nil, err
	}

	claims, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims format")
	}

	user := NewSysUser()
	user.Id = int(claims["user_id"].(float64))
	user.Nickname = claims["nickName"].(string)
	user.RoleIds = claims["role"].(string)

	return user, nil
}
