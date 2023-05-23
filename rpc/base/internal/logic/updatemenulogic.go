package logic

import (
	"bolg/rpc/utils"
	"context"

	"bolg/rpc/base/base_client"
	"bolg/rpc/base/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuLogic) UpdateMenu(in *base_client.UpdateMenuRequest) (*base_client.MenuListResponse, error) {
	menuData, err := utils.MenuRpcToGorm(in.GetMenuData())

	if err != nil {
		l.Errorf("UpdateMenu->MenuRpcToGorm;err=%v", err)
		return nil, err
	}
	err = l.svcCtx.DbEngin.Save(&menuData).Error
	if err != nil {
		l.Errorf("UpdateMenu->l.svcCtx.DbEngin.Save;err=%v", err)
		return nil, err
	}

	responseData, err := utils.MenuGormToRpc(menuData)
	if err != nil {
		l.Errorf("UpdateMenu->MenuGormToRpc;err=%v", err)
		return nil, err
	}

	return &base_client.MenuListResponse{
		MenuData: responseData,
	}, nil
}