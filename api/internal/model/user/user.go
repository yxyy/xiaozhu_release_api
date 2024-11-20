package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	common2 "xiaozhu/api/internal/logic/common"
	"xiaozhu/api/utils"
)

type User struct {
	common2.Model
	Account       string `json:"account" form:"account" gorm:""`
	Password      string `json:"-" form:"password" gorm:"password"`
	Salt          string `json:"-" gorm:"salt"`
	Nickname      string `json:"nickname" form:"nickname" gorm:"nickname"`
	Phone         int64  `json:"phone" form:"phone" gorm:"phone"`
	Wechat        string `json:"wechat" form:"wechat" gorm:"wechat"`
	Email         string `json:"email" form:"email" gorm:"email"`
	RoleIds       int    `json:"group_id" form:"group_id" gorm:"group_id"`
	Avatar        string `json:"avatar" form:"avatar" gorm:"avatar"`
	Status        int64  `json:"status" form:"status" gorm:"status"`
	LastLoginIp   string `json:"last_login_ip" form:"last_login_ip" gorm:"last_login_ip"`
	LastLoginTime int64  `json:"last_login_time" form:"last_login_time" gorm:"last_login_time"`
	Remarks       string `json:"remarks" form:"remarks" gorm:"remarks"`
}

type SysUser struct {
	common2.Model
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

type TokenInfo struct {
	SysUser
	GroupName    string `json:"group_name"`
	PermissionId int    `json:"permission_id"`
}

func NewSysUser() *SysUser {
	return &SysUser{}
}

func (u *SysUser) UserInfo() (user *User, err error) {
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

func (u *SysUser) List(ctx context.Context, params common2.Params) (user []*SysUser, total int64, err error) {
	tx := utils.MysqlDb.Model(&u).WithContext(ctx)
	if u.Id > 0 {
		tx = tx.Where("id", u.Id)
	}
	if u.DeptId != nil {
		tx = tx.Where("dept_id", u.DeptId)
	}
	if u.Status != nil && *u.Status != "" {
		statusInt, err := strconv.Atoi(*u.Status)
		if err != nil {
			return nil, 0, err
		}
		tx = tx.Where("status", statusInt)
	}
	if u.Mobile != "" {
		tx = tx.Where("mobile", u.Mobile)
	}
	if u.Account != "" {
		tx = tx.Where("account like ? ", "%"+u.Account+"%")
	}
	if u.Email != "" {
		tx = tx.Where("email", u.Email)
	}
	if u.Wechat != "" {
		tx = tx.Where("wechat", u.Wechat)
	}
	if u.Nickname != "" {
		tx = tx.Where("nickname like ?", "%"+u.Nickname+"%")
	}

	if err = tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err = tx.Offset(params.Offset).Limit(params.Limit).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	return
}

func (u *SysUser) All() (users []*User, err error) {

	if err = utils.MysqlDb.Model(&u).Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *SysUser) Create(ctx context.Context) error {

	return utils.MysqlDb.Model(&u).WithContext(ctx).Create(&u).Error
}

func (u *SysUser) Update(ctx context.Context) error {

	return utils.MysqlDb.Model(&u).WithContext(ctx).Where("id", u.Id).Updates(&u).Error
}

func (u *SysUser) Remove() error {
	return utils.MysqlDb.Model(&u).Where("id", u.Id).Delete(&u).Error
}

func (u *SysUser) Get(ctx context.Context) (user SysUser, err error) {

	tx := utils.MysqlDb.Model(&u).WithContext(ctx)
	if u.Id <= 0 {
		tx = tx.Where("id", u.Id)
	}
	if u.Account != "" {
		tx = tx.Where("account", u.Account)
	}
	if err = tx.Find(&user).Error; err != nil {
		return
	}
	return
}

func (p *SysUser) GetAll(ctx context.Context) (list []*common2.IdName, err error) {
	err = utils.MysqlDb.Model(&p).WithContext(ctx).Select("id,nickname as name").Scan(&list).Error
	return
}
