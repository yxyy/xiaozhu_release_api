package assets

import (
	"context"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"xiaozhu/backend/internal/logic/system"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/internal/model/market/assets"
	"xiaozhu/backend/utils"
)

type AccountLogic struct {
	ctx context.Context
	assets.Account
	common.Params
	FilePath string `json:"filepath"`
}

// type OtherParams struct {
// 	common.Params
// 	ChannelId int `json:"channel_id"`
// }

type AccountListResponse struct {
	List  []*assets.AdAccountList `json:"list"`
	Total int64                   `json:"total"`
}

func NewAccountLogic(ctx context.Context) *AccountLogic {
	return &AccountLogic{ctx: ctx}
}

func (l *AccountLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *AccountLogic) Create() error {
	if l.Name == "" {
		return errors.New("名称不能为空")
	}
	if l.ProjectId < 1 {
		return errors.New("项目不能为空")
	}

	l.Lot = utils.Uuid()
	l.OptUser = l.ctx.Value("userId").(int)

	return l.Account.Create(l.ctx)
}

func (l *AccountLogic) BatchCreate() error {
	if l.FilePath == "" {
		return errors.New("附件不能为空")
	}
	if l.ProjectId < 1 {
		return errors.New("项目不能为空")
	}

	f, err := excelize.OpenFile(l.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	var list []*assets.Account

	lot := utils.Uuid()
	for k, row := range rows {
		if k == 0 {
			continue
		}
		if row[0] == "" {
			return errors.New(fmt.Sprintf("第%d行，广告主Id不能为空", k))
		}
		if row[1] == "" {
			return errors.New(fmt.Sprintf("第%d行，广告主名称不能为空", k))
		}

		shortName := row[1]
		if row[2] != "" {
			shortName = row[2]
		}

		Uid, err := strconv.Atoi(row[0])
		if err != nil {
			return errors.New(fmt.Sprintf("第%d行，无效发广告主Id:%s不能为空:"+err.Error(), k, row[0]))
		}
		tmp := &assets.Account{
			Name:      row[1],
			ShortName: shortName,
			ProjectId: l.ProjectId,
			Uid:       Uid,
			Owner:     l.ctx.Value("userId").(int), // TODO 后期中文名转对应id
			Lot:       lot,
			Model: common.Model{
				OptUser: l.ctx.Value("userId").(int),
			},
		}

		list = append(list, tmp)
	}

	if err = l.Account.BatchCreate(l.ctx, list); err != nil {
		return err
	}

	return os.Remove(l.FilePath)
}

func (l *AccountLogic) Update() error {
	if l.Id <= 0 {
		return errors.New("id无效")
	}
	l.OptUser = l.ctx.Value("userId").(int)

	return l.Account.Update(l.ctx)
}

func (l *AccountLogic) List() (resp *AccountListResponse, err error) {

	list, total, err := l.Account.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}

	resp = new(AccountListResponse)

	userLogic := system.NewUserLogic(l.ctx)
	userInfo, err := userLogic.ListAll()
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		v.OwnerName = userInfo[v.Owner].Name
		v.OptUserName = userInfo[v.OptUser].Name
		resp.List = append(resp.List, v)
	}

	resp.Total = total

	return
}

func (l *AccountLogic) ListAll() (list map[int]*common.IdName, err error) {

	resp, err := l.Account.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}
	return utils.ConvertIdNameMapById(resp), nil
}
