package user

import (
	"context"
	"errors"
	"xiaozhu/internal/logic/common"
	"xiaozhu/internal/model/user"
)

type MemberLogic struct {
	ctx context.Context
	common.RequestForm
}

type MemberInfoResponse struct {
	UserId   int    `json:"user_id"`
	Account  string `json:"account"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Balance  int    `json:"balance"`
}

func NewMemberLogic(ctx context.Context) *MemberLogic {
	return &MemberLogic{ctx: ctx}
}

func (l *MemberLogic) Info() (*MemberInfoResponse, error) {

	userId, ok := l.ctx.Value("userId").(int)
	if userId == 0 || !ok {
		return nil, errors.New("请先登录")
	}

	member := user.NewMemberInfo()
	member.Id = userId
	err := member.MemberInfo(l.ctx)
	if err != nil {
		return nil, err
	}

	resp := new(MemberInfoResponse)
	resp.UserId = member.UserId
	resp.Account = member.Account
	resp.NickName = member.Nickname
	resp.Avatar = member.Avatar
	resp.Balance = member.Balance

	return resp, nil
}
