package assets

import (
	"errors"
	"xiaozhu/internal/logic/conmon"
	"xiaozhu/internal/mapping"
	"xiaozhu/internal/model/assets"
	"xiaozhu/internal/model/common"
)

type ServiceChannel struct {
	assets.Channel
	conmon.Format
}

func NewServiceChannel() ServiceChannel {
	return ServiceChannel{}
}

func (c ServiceChannel) Create() error {
	if c.Name == "" {
		return errors.New("名称不能为空")
	}
	if c.Flag == "" {
		return errors.New("标识不能为空")
	}

	return c.Channel.Create()
}

func (c ServiceChannel) Update() error {
	if c.Id <= 0 {
		return errors.New("id无效")
	}

	return c.Channel.Update()
}

func (c ServiceChannel) List(params common.Params) (list []*ServiceChannel, total int64, err error) {
	params.Verify()
	channels, total, err := c.Channel.List(params)
	if err != nil {
		return nil, 0, err
	}

	userMap, err := mapping.User()
	if err != nil {
		return nil, 0, err
	}

	for _, v := range channels {
		format := conmon.Formats(v.Model)
		format.OptUserName = userMap[v.OptUser]
		node := &ServiceChannel{
			Format:  format,
			Channel: *v,
		}
		list = append(list, node)
	}

	return
}

func (c ServiceChannel) Lists() (list []*assets.Channel, err error) {

	return c.GetAll()
}
