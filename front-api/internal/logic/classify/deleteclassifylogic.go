package classify

import (
	"bolg/rpc/classify/classify_client"
	"context"
	"fmt"

	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassifyLogic {
	return &DeleteClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClassifyLogic) DeleteClassify(req *types.DeleteClassifyReq) (resp *types.ClassifyRespose, err error) {
	classify, err := l.svcCtx.Classify.DeleteClassify(l.ctx, &classify_client.DeleteClassifyRequest{
		Id: []uint64{req.ID},
	})
	if err != nil {
		return &types.ClassifyRespose{
			Code:    200,
			Data:    nil,
			Message: "NULL DATA",
		}, nil
	}

	fmt.Println("classifyclassify", classify)

	return &types.ClassifyRespose{
		Code:    200,
		Data:    nil,
		Message: "SUCCESS",
	}, nil
}
