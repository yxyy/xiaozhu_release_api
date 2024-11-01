package assets

import (
	"errors"
	"xiaozhu/internal/logic/conmon"
	"xiaozhu/internal/mapping"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
)

type ServiceAppType struct {
	assets.AppType
	conmon.Format
	CompanyName string `json:"company_name"`
	TypeName    string `json:"type_name"`
}

func NewServiceAppType() ServiceAppType {
	return ServiceAppType{}
}

func (c ServiceAppType) List(params common.Params) (sc []*ServiceAppType, total int64, err error) {
	params.Verify()
	list, total, err := c.AppType.List(params)
	if err != nil {
		return nil, 0, err
	}
	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}

	for _, v := range list {

		format := conmon.Formats(v.Model)
		format.OptUserName = userMap[v.OptUser]
		serviceAppType := &ServiceAppType{
			AppType: *v,
			Format:  format,
		}

		sc = append(sc, serviceAppType)
	}

	return
}

func (c ServiceAppType) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}

	return c.AppType.Create()
}

func (c ServiceAppType) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.AppType.Update()
}

func (c ServiceAppType) Lists() (sc []*assets.AppType, err error) {

	return c.AppType.GetAll()
}
