package user

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"xiaozhu/internal/config"
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
	ProfileId     int    `json:"profile_id" gorm:"column:id"`
	UserId        int    `json:"user_id" gorm:"user_id"`   // 用户ID
	Nickname      string `json:"nickname" gorm:"nickname"` // 昵称
	Mobile        string `json:"mobile" gorm:"mobile"`     // 绑定的手机号
	Wechat        string `json:"wechat" form:"wechat" gorm:"wechat"`
	TradePassword string `json:"trade_password" gorm:"trade_password"`   // 安全码
	Balance       int    `json:"balance" gorm:"balance"`                 // 余额
	FullName      string `json:"full_name" gorm:"full_name"`             // 姓名
	RegTime       int64  `json:"reg_time" gorm:"reg_time"`               // 注册时间
	RegIp         string `json:"reg_ip" gorm:"reg_ip"`                   // 注册IP
	LastLoginTime int64  `json:"last_login_time" gorm:"last_login_time"` // 最后登陆时间
	LastLoginIp   string `json:"last_ip" gorm:"last_ip"`                 // 最后登陆IP
	LastLoginWay  int8   `json:"last_login_way" gorm:"last_login_way"`   // 最后登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)
	AreaCode      string `json:"area_code" gorm:"area_code"`             // 注册地区码
	Area          string `json:"area" gorm:"area"`                       // 注册地区
	GameId        int    `json:"game_id" gorm:"game_id"`                 // 游戏ID
	Email         string `json:"email" gorm:"email"`                     // 用户邮箱
	Sex           int8   `json:"sex" gorm:"sex"`                         // 2:未知 1:男 0: 女
	Avatar        string `json:"avatar" gorm:"avatar"`                   // 头像
	Vip           int16  `json:"vip" gorm:"vip"`                         // 会员VIP等级
	DeviceId      string `json:"device_id" gorm:"device_id"`             // 设备号
	SysModel      string `json:"sys_model" gorm:"sys_model"`             // 机型
	LoginTimes    int32  `json:"login_times" gorm:"login_times"`         // 登陆次数
	Remark        string `json:"remark" gorm:"remark"`                   // 备注
	// Integral      int    `json:"integral" gorm:"integral"`               // 积分
	CreatedAt int64 `json:"created_at,omitempty" form:"created_at" gorm:"created_at"`
	UpdatedAt int64 `json:"updated_at,omitempty" form:"updated_at" gorm:"updated_at"`
	// DeletedAt int64 `json:"deleted_at,omitempty" form:"deleted_at" gorm:"deleted_at"`
}

type MemberInfo struct {
	Member
	MemberProfile
}

func NewMemberInfo() *MemberInfo {
	return &MemberInfo{}
}

func (i *MemberInfo) TableName() string {
	return "member"
}

func (i *MemberInfo) MemberInfo(ctx context.Context) (err error) {
	if i.Id <= 0 && i.Account == "" && i.Email == "" {
		return errors.New("参数错误")
	}
	tx := config.MysqlDefaultDb.
		Table("member").
		Select("member.*", "member_profile.*", "member_profile.id as profile_id").
		WithContext(ctx).
		Joins("left join member_profile on member_profile.user_id = member.id")

	if i.Id > 0 {
		tx = tx.Where("member.id", i.Id)
	}
	if i.Account != "" {
		tx = tx.Where("account", i.Account)
	}
	if i.Email != "" {
		tx = tx.Where("email", i.Email)
	}
	if err = tx.First(&i).Error; err != nil {
		return err
	}

	return
}

func (i *MemberInfo) Create(ctx context.Context) error {
	return config.MysqlDefaultDb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 插入 member 表
		if err := tx.Table("member").Create(&i.Member).Error; err != nil {
			return fmt.Errorf("插入 member 表失败: %w", err)
		}
		if i.Member.Id == 0 {
			return errors.New("member 插入失败")
		}

		i.MemberProfile.UserId = i.Member.Id

		// 插入 member_profile 表
		if err := tx.Table("member_profile").Create(&i.MemberProfile).Error; err != nil {
			return fmt.Errorf("插入 member_profile 表失败: %w", err)
		}

		return nil
	})
}

func (u *Member) Update(ctx context.Context) error {

	return config.MysqlDefaultDb.Model(&u).WithContext(ctx).Where("id", u.Id).Updates(&u).Error
}
