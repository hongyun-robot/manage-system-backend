package classify

import (
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"bolg/front-api/internal/utils"
	"bolg/rpc/classify/classify_client"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassifyLogic {
	return &UpdateClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClassifyLogic) UpdateClassify(req *types.UpdateClassifyReq) (resp *types.ClassifyRespose, err error) {
	classify, err := l.svcCtx.Classify.UpdateClassify(l.ctx, &classify_client.UpdateClassifyRequest{
		Id:         req.ID,
		Name:       req.Name,
		Icon:       req.Icon,
		ArticleIds: req.ArticlesIds,
	})

	if err != nil {
		return &types.ClassifyRespose{
			Code:    200,
			Data:    nil,
			Message: "ERROR",
		}, nil
	}

	classifyData := utils.ClassifyConversionType(&classify.Classify)

	return &types.ClassifyRespose{
		Code:    200,
		Data:    classifyData,
		Message: "SUCCESS",
	}, nil
}
