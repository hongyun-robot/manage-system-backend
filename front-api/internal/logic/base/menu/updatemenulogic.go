package menu

import (
	"bolg/front-api/global"
	"bolg/front-api/internal/utils"
	"bolg/rpc/base/base_client"
	"context"

	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (resp *types.MenuRespose, err error) {
	menuData, err := utils.MenuApiToRpc(req.Data)
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.CONVERSION,
		}, nil
	}

	updateMenu, err := l.svcCtx.Base.UpdateMenu(l.ctx, &base_client.UpdateMenuRequest{MenuData: menuData})
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.ERROR,
		}, nil
	}

	responseData, err := utils.MenuRpcToApi(updateMenu.GetMenuData())
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.CONVERSION,
		}, nil
	}

	return &types.MenuRespose{
		Code:    200,
		Data:    responseData,
		Message: global.SUCCESS,
	}, nil
}