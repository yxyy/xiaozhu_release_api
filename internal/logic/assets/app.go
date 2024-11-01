package assets

import (
	"errors"
	"xiaozhu/internal/logic/conmon"
	"xiaozhu/internal/mapping"
	"xiaozhu/internal/model/common"
)

type ServiceApp struct {
	assets.App
	conmon.Format
	CompanyName string `json:"company_name"`
	TypeName    string `json:"type_name"`
}

func NewServiceApp() ServiceApp {
	return ServiceApp{}
}

func (c ServiceApp) List(params common.Params) (sc []*ServiceApp, total int64, err error) {
	params.Verify()
	companys, total, err := c.App.List(params)
	if err != nil {
		return nil, 0, err
	}
	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}

	companies, err := mapping.Company()
	if err != nil {
		return nil, 0, err
	}

	appType, err := mapping.AppType()
	if err != nil {
		return nil, 0, err
	}

	for _, v := range companys {

		format := conmon.Formats(v.Model)
		format.OptUserName = userMap[v.OptUser]
		serviceApp := &ServiceApp{
			App:         *v,
			Format:      format,
			CompanyName: companies[v.CompanyId],
			TypeName:    appType[v.Type],
		}

		sc = append(sc, serviceApp)
	}

	return
}

func (c ServiceApp) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}

	if c.CompanyId <= 0 {
		return errors.New("研发公司不能为空")
	}

	return c.App.Create()
}

func (c ServiceApp) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.App.Update()
}

func (c ServiceApp) Lists() (sc []*assets.App, err error) {

	return c.App.GetAll()
}
