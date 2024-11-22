package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"xiaozhu/api/utils"
)

type Member struct {
	Id          int    `json:"id"`
	PromoteCode string `json:"promote_code"`
	AccountType int    `json:"account_type"`
	Account     string `json:"account" form:"account" gorm:""`
	Password    string `json:"-" form:"password" gorm:"password"`
	Salt        string `json:"-" gorm:"salt"`
	Status      int    `json:"status"`
}

type MemberProfile struct {
	Id            int64  `json:"id" gorm:"id"`
	UserId        int64  `json:"user_id" gorm:"user_id"`               // 用户ID
	Nickname      string `json:"nickname" gorm:"nickname"`             // 昵称
	Mobile        string `json:"mobile" gorm:"mobile"`                 // 绑定的手机号
	TradePassword string `json:"trade_password" gorm:"trade_password"` // 安全码
	Balance       int64  `json:"balance" gorm:"balance"`               // 余额
	FullName      string `json:"full_name" gorm:"full_name"`           // 姓名
	RegTime       int64  `json:"reg_time" gorm:"reg_time"`             // 注册时间
	RegIp         string `json:"reg_ip" gorm:"reg_ip"`                 // 注册IP
	LastTime      int64  `json:"last_time" gorm:"last_time"`           // 最后登陆时间
	LastIp        string `json:"last_ip" gorm:"last_ip"`               // 最后登陆IP
	LastLoginWay  int8   `json:"last_login_way" gorm:"last_login_way"` // 最后登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)
	AreaCode      string `json:"area_code" gorm:"area_code"`           // 注册地区码
	Area          string `json:"area" gorm:"area"`                     // 注册地区
	GameId        int64  `json:"game_id" gorm:"game_id"`               // 游戏ID
	Qq            string `json:"qq" gorm:"qq"`
	Email         string `json:"email" gorm:"email"`             // 用户邮箱
	Sex           int8   `json:"sex" gorm:"sex"`                 // 2:未知 1:男 0: 女
	Avatar        string `json:"avatar" gorm:"avatar"`           // 头像
	Vip           int16  `json:"vip" gorm:"vip"`                 // 会员VIP等级
	DeviceId      string `json:"device_id" gorm:"device_id"`     // 设备号
	SysModel      string `json:"sys_model" gorm:"sys_model"`     // 机型
	LoginTimes    int32  `json:"login_times" gorm:"login_times"` // 登陆次数
	Remark        string `json:"remark" gorm:"remark"`           // 备注
	Integral      int64  `json:"integral" gorm:"integral"`       // 积分
}

func NewMember() *Member {
	return &Member{}
}

func (u *Member) UserInfo() (user *Member, err error) {
	if u.Id <= 0 && u.Account == "" {
		return user, errors.New("参数错误")
	}
	tx := utils.MysqlDb.Model(&u)
	if u.Id > 0 {
		tx = tx.Where("id", u.Id)
	}
	if u.Account != "" {
		tx = tx.Where("account", u.Account)
	}
	if err = tx.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return
}

func (u *Member) Create(ctx context.Context) error {

	return utils.MysqlDb.Model(&u).WithContext(ctx).Create(&u).Error
}

func (u *Member) Update(ctx context.Context) error {

	return utils.MysqlDb.Model(&u).WithContext(ctx).Where("id", u.Id).Updates(&u).Error
}
