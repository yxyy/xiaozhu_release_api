package assets

import (
	"context"
	"errors"
	"xiaozhu/backend/internal/model/assets"
	"xiaozhu/backend/internal/model/common"
	"xiaozhu/backend/utils"
)

type AppLogic struct {
	ctx    context.Context
	App    assets.App
	Params common.Params
}

type AppListResponse struct {
	List  []*assets.App `json:"list"`
	Total int64         `json:"total"`
}

func NewAppLogic(ctx context.Context) *AppLogic {
	return &AppLogic{ctx: ctx}
}

func (l *AppLogic) GetParams() *common.Params {
	l.Params.Verify()
	return &l.Params
}

func (l *AppLogic) List() (reps *AppListResponse, err error) {

	list, total, err := l.App.List(l.ctx, l.GetParams())
	if err != nil {
		return nil, err
	}
	// userMap, err := mapping.User()
	// if err != nil {
	// 	return nil, 0, err
	// }

	// companies, err := mapping.Company()
	// if err != nil {
	// 	return nil, 0, err
	// }

	// appType, err := mapping.AppType()
	// if err != nil {
	// 	return nil, 0, err
	// }

	// for _, v := range app {
	//
	// 	format := conmon.Formats(v.Model)
	// 	format.OptUserName = userMap[v.OptUser]
	// 	serviceApp := &AppLogic{
	// 		App:         *v,
	// 		Format:      format,
	// 		CompanyName: companies[v.CompanyId],
	// 		TypeName:    appType[v.GameClass],
	// 	}
	//
	// 	sc = append(sc, serviceApp)
	// }

	reps = new(AppListResponse)
	reps.List = list
	reps.Total = total

	return
}

func (l *AppLogic) Create() error {
	if l.App.AppName == "" {
		return errors.New("名称不能为空")
	}

	if l.App.CompanyId <= 0 {
		return errors.New("研发公司不能为空")
	}

	return l.App.Create(l.ctx)
}

func (l *AppLogic) Update() error {
	if l.App.Id <= 0 {
		return errors.New("id无效")
	}

	return l.App.Update(l.ctx)
}

func (l *AppLogic) ListAll() (resp map[int]*common.IdName, err error) {

	app, err := l.App.GetAll(l.ctx)
	if err != nil {
		return nil, err
	}

	return utils.ConvertIdNameMapById(app), nil
}
