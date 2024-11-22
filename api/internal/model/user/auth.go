package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
	"xiaozhu/api/internal/logic/common"
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

type SysUser struct {
	common.Model
	// Id           uint64 `json:"id" gorm:"id"`
	Account      string  `json:"account" gorm:"account"`     // 帐户名
	Password     string  `json:"password,-" gorm:"password"` // 登陆密码
	Salt         string  `json:"-" gorm:"salt"`              // 盐值
	Status       *string `json:"status" gorm:"status"`       // -1:未激活，0：正常，1：禁用
	Nickname     string  `json:"nickname" gorm:"nickname"`   // 昵称
	Wechat       string  `json:"wechat" form:"wechat" gorm:"wechat"`
	Mobile       string  `json:"mobile" gorm:"mobile"`                 // 绑定的手机号
	FullName     string  `json:"full_name" gorm:"full_name"`           // 姓名
	RegTime      int64   `json:"reg_time" gorm:"reg_time"`             // 注册时间
	RegIp        string  `json:"reg_ip" gorm:"reg_ip"`                 // 注册IP
	LastTime     int64   `json:"last_time" gorm:"last_time"`           // 最后登陆时间
	LastIp       string  `json:"last_ip" gorm:"last_ip"`               // 最后登陆IP
	LastLoginWay int8    `json:"last_login_way" gorm:"last_login_way"` // 最后登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)
	RoleIds      string  `json:"role_ids" gorm:"role_ids"`
	DeptId       *int    `json:"deptId" gorm:"dept_id"`    // 部门ID
	PostIds      string  `json:"post_ids" gorm:"post_ids"` // 岗位编号数组
	Qq           string  `json:"qq" gorm:"qq"`
	Email        string  `json:"email" gorm:"email"`             // 用户邮箱
	Sex          int8    `json:"sex" gorm:"sex"`                 // 2:未知 1:男 0: 女
	Avatar       string  `json:"avatar" gorm:"avatar"`           // 头像
	LoginTimes   int32   `json:"login_times" gorm:"login_times"` // 登陆次数
	Remark       string  `json:"remark" gorm:"remark"`           // 备注
}

func NewSysUser() *SysUser {
	return &SysUser{}
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
