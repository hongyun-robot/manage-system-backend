package logic

import (
	"bolg/rpc/base/base_client"
	"bolg/rpc/base/internal/svc"
	"bolg/rpc/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMenuLogic) AddMenu(in *base_client.MenuListResponse) (*base_client.MenuListResponse, error) {
	data, err := utils.MenuRpcToGorm(in.GetMenuData())
	if err != nil {
		l.Errorf("AddMenu->MenuRpcToGorm;err=%v", err)
		return nil, err
	}
	_ = l.svcCtx.DbEngin.Create(&data)
	responseData, err := utils.MenuGormToRpc(data)
	if err != nil {
		l.Errorf("AddMenu->MenuGormToRpc;err=%v", err)
		return nil, err
	}
	return &base_client.MenuListResponse{
		MenuData: responseData,
	}, nil
}
