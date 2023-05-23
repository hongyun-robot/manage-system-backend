package menu

import (
	"bolg/front-api/global"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"bolg/front-api/internal/utils"
	"bolg/rpc/base/base_client"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuLogic) GetMenu(req *types.GetMenutReq) (resp *types.MenuRespose, err error) {
	fmt.Println(req)
	menuListResponse, err := l.svcCtx.Base.GetMenu(l.ctx, &base_client.GetMenuRequest{})
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.ERROR,
		}, nil
	}
	responseData, err := utils.MenuRpcToApi(menuListResponse.GetMenuData())
	if err != nil {
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.ERROR,
		}, nil
	}

	return &types.MenuRespose{
		Code:    200,
		Data:    responseData,
		Message: global.SUCCESS,
	}, nil
}
