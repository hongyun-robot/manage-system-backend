package menu

import (
	"bolg/front-api/global"
	"bolg/rpc/base/base_client"
	"context"
	"fmt"

	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (resp *types.MenuRespose, err error) {
	_, err = l.svcCtx.Base.DeleteMenu(l.ctx, &base_client.DeleteMenuRequest{Id: []uint64{req.ID}})
	if err != nil {
		fmt.Println("err", err)
		return &types.MenuRespose{
			Code:    500,
			Data:    nil,
			Message: global.ERROR,
		}, nil
	}

	return &types.MenuRespose{
		Code:    200,
		Data:    nil,
		Message: global.SUCCESS,
	}, nil
}
