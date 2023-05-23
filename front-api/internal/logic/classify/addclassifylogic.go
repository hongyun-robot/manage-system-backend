package classify

import (
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"bolg/rpc/classify/classify_client"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassifyLogic {
	return &AddClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassifyLogic) AddClassify(req *types.AddClassifyReq) (resp *types.ClassifyRespose, err error) {
	fmt.Println("reqreqreq", req)
	responseData, err := l.svcCtx.Classify.AddClassify(l.ctx, &classify_client.AddClassifyRequest{
		Name:       req.Name,
		Icon:       req.Icon,
		ArticleIds: req.ArticleIds,
	})
	if err != nil {
		return &types.ClassifyRespose{
			Code:    500,
			Data:    nil,
			Message: "error",
		}, err
	}
	responseClassify := responseData.Classify[0]
	articleData := make([]*types.ClassifyArticleData, 0, len(responseClassify.ArticleData))
	for _, item := range responseClassify.ArticleData {
		articleData = append(articleData, &types.ClassifyArticleData{
			ID:       item.Id,
			Title:    item.Title,
			Content:  item.Content,
			UpdateAt: item.UpdateAt,
			Status:   item.Status,
			Draft:    item.Draft,
		})
	}
	data := []*types.ClassifyData{
		{
			ID:          responseClassify.Id,
			Name:        responseClassify.Name,
			CreateAt:    responseClassify.CreateAt,
			UpdateAt:    responseClassify.UpdateAt,
			Icon:        responseClassify.Icon,
			ArticleData: articleData,
		},
	}

	return &types.ClassifyRespose{
		Code:    200,
		Data:    data,
		Message: "success",
	}, nil
}
