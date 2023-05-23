package menu

import (
	"bolg/front-api/global"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"bolg/front-api/internal/utils"
	"bolg/rpc/base/base_client"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuLogic) AddMenu(req *types.AddMenuReq) (resp *types.MenuRespose, err error) {
	MenuData, err := utils.MenuApiToRpc(req.Data)
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.CONVERSION,
		}, nil
	}
	menu, err := l.svcCtx.Base.AddMenu(l.ctx, &base_client.MenuListResponse{MenuData: MenuData})
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.ERROR,
		}, nil
	}

	responseData, err := utils.MenuRpcToApi(menu.GetMenuData())
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
