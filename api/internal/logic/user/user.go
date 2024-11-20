package user

import (
	"context"
	"errors"
	"strconv"
	common2 "xiaozhu-api/internal/logic/common"
	"xiaozhu-api/internal/model/user"
	"xiaozhu-api/utils"
)

type UserLogic struct {
	ctx     context.Context
	SysUser user.SysUser
	common2.Params
}

func NewUserLogic(ctx context.Context) UserLogic {
	return UserLogic{ctx: ctx}
}

type Dept struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type SysUserResponse struct {
	Id           int    `json:"id"`
	Account      string `json:"account"`   // 帐户名
	Status       int8   `json:"status"`    // -1:未激活，0：正常，1：禁用
	Nickname     string `json:"nickname" ` // 昵称
	Wechat       string `json:"wechat"`
	Mobile       string `json:"mobile"`         // 绑定的手机号
	FullName     string `json:"full_name"`      // 姓名
	RegTime      int64  `json:"reg_time"`       // 注册时间
	RegIp        string `json:"reg_ip"`         // 注册IP
	LastTime     int64  `json:"last_time"`      // 最后登陆时间
	LastIp       string `json:"last_ip"`        // 最后登陆IP
	LastLoginWay int8   `json:"last_login_way"` // 最后登陆方式： 0(visitor)、1(email)、 2(facebook)、3(google)、4(apple)
	RoleIds      string `json:"roleIds"`
	Dept         Dept   `json:"dept"`
	Qq           string `json:"qq"`
	Email        string `json:"email"`       // 用户邮箱
	Sex          int8   `json:"sex"`         // 2:未知 1:男 0: 女
	Avatar       string `json:"avatar"`      // 头像
	LoginTimes   int32  `json:"login_times"` // 登陆次数
	Remark       string `json:"remark"`      // 备注
	CreatedAt    int64  `json:"createTime,omitempty"`
}

type UserListResponse struct {
	List  []*SysUserResponse `json:"list"`
	Total int64              `json:"total"`
}

func (l *UserLogic) ListLogic() (resp *UserListResponse, err error) {

	list, total, err := l.SysUser.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(UserListResponse)
	for _, v := range list {
		status, err := strconv.Atoi(*v.Status)
		if err != nil {
			return nil, err
		}

		tmp := &SysUserResponse{
			Id:           v.Id,
			Account:      v.Account,
			Status:       int8(status),
			Nickname:     v.Nickname,
			Wechat:       v.Wechat,
			Mobile:       v.Mobile,
			FullName:     v.FullName,
			RegTime:      v.RegTime,
			RegIp:        v.RegIp,
			LastTime:     v.LastTime,
			LastIp:       v.LastIp,
			LastLoginWay: v.LastLoginWay,
			RoleIds:      v.RoleIds,
			Dept: Dept{
				Id:   *v.DeptId,
				Name: "6666",
			},
			Qq:         v.Qq,
			Email:      v.Email,
			Sex:        v.Sex,
			Avatar:     v.Avatar,
			LoginTimes: v.LoginTimes,
			Remark:     v.Remark,
		}

		resp.List = append(resp.List, tmp)
	}

	resp.Total = total

	return
}

func (l *UserLogic) GetParams() common2.Params {
	l.Params.Verify()
	return l.Params
}

func (l UserLogic) UserInfos() (map[string]interface{}, error) {
	// fmt.Printf("%#v\n", l)
	// user, err := l.UserInfo()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// servicesMenu := NewServicesMenu()
	// getTree, err := servicesMenu.GetTree()
	// if err != nil {
	// 	return nil, err
	// }
	//
	// data := make(map[string]interface{})
	// data["userInfo"] = user
	// data["menus"] = getTree

	return nil, nil
}

func (l *UserLogic) Create() error {
	if l.SysUser.Account == "" {
		return errors.New("账号不能为空")
	}
	if l.SysUser.Password == "" {
		return errors.New("密码不能为空")
	}
	if len(l.SysUser.Password) < 6 {
		return errors.New("密码长度不能小于6位")
	}

	user, err := l.SysUser.Get(l.ctx)
	if err != nil {
		return err
	}
	if user.Account != "" {
		return errors.New("账号已存在")
	}

	if l.SysUser.Nickname == "" {
		l.SysUser.Nickname = l.SysUser.Account
	}
	l.SysUser.Salt = utils.Salt()
	l.SysUser.Password = utils.Md5SaltAndPassword(l.SysUser.Salt, l.SysUser.Password)
	// l.SysUser.CreatedAt = time.Now().Unix()
	// l.SysUser.UpdatedAt = time.Now().Unix()

	return l.SysUser.Create(l.ctx)
}

func (l *UserLogic) Update() error {

	if l.SysUser.Id <= 0 {
		return errors.New("Id无效")
	}
	// if l.SysUser.Status < 0 || l.SysUser.Status > 2 {
	// 	return errors.New("状态无效")
	// }

	user := user.SysUser{
		Nickname: l.SysUser.Nickname,
		Status:   l.SysUser.Status,
		RoleIds:  l.SysUser.RoleIds,
		Avatar:   l.SysUser.Avatar,
		LastIp:   l.SysUser.LastIp,
		LastTime: l.SysUser.LastTime,
		Remark:   l.SysUser.Remark,
		Model: common2.Model{
			Id:      l.SysUser.Id,
			OptUser: l.ctx.Value("userId").(int),
		},
	}
	l.SysUser = user

	return l.SysUser.Update(l.ctx)
}

func (l *UserLogic) SaveRole() error {
	if l.SysUser.Id < 1 {
		return errors.New("id不能为空")
	}
	if l.SysUser.RoleIds == "" {
		return errors.New("角色不能为空")
	}

	user := user.SysUser{
		RoleIds: l.SysUser.RoleIds,
		Model: common2.Model{
			Id:      l.SysUser.Id,
			OptUser: l.ctx.Value("userId").(int),
		},
	}
	l.SysUser = user

	return l.SysUser.Update(l.ctx)
}

func (l *UserLogic) Remove() error {
	if l.SysUser.Id <= 0 {
		return errors.New("Id无效")
	}

	return l.SysUser.Remove()
}

func (l *UserLogic) ListAll() (list map[int]*common2.IdName, err error) {

	resp, err := l.SysUser.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}
