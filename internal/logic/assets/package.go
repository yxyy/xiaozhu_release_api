package assets

import (
	"errors"
	"xiaozhu/internal/logic/conmon"
	"xiaozhu/internal/model/common"
	"xiaozhu/utils"
)

type ServicePackage struct {
	assets.Package
	conmon.Format
}

func NewServicePackage() ServicePackage {
	return ServicePackage{}
}

func (p ServicePackage) List(params common.Params) (list []map[string]interface{}, total int64, err error) {
	params.Verify()
	// list, total, err = p.Package.ListLogic(params)
	return p.Package.List(params)
}

func (p ServicePackage) Create() error {
	if p.Name == "" {
		return errors.New("渠道包名称不能为空")
	}

	if p.ChannelId <= 0 {
		return errors.New("渠道不能为空")
	}

	if p.GameId <= 0 {
		return errors.New("游戏不能为空")
	}

	if p.Status < 0 {
		return errors.New("状态无效")
	}

	p.Campaign = utils.Random(8)
	if p.Campaign == "" {
		return errors.New("渠道包标识生成错误")
	}

	return p.Package.Create()
}

func (p ServicePackage) Update() error {

	if p.Id <= 0 {
		return errors.New("渠道不能为空")
	}

	if p.Status < 0 {
		return errors.New("状态无效")
	}

	return p.Package.Update()
}
