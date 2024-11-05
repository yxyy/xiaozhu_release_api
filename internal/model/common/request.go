package common

import (
	"fmt"
)

type Params struct {
	Page   int `json:"page,omitempty" form:"page"`
	Limit  int `json:"limit,omitempty" form:"limit"`
	Offset int `json:"offset,omitempty" form:"offset"`
}

func NewParams() Params {
	return Params{
		Page:  1,
		Limit: 10,
	}
}

func (p *Params) Verify() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	p.Offset = (p.Page - 1) * p.Limit

	fmt.Println(p.Page, p.Limit, p.Offset)
}
