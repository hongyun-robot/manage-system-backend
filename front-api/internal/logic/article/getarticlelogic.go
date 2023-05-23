package article

import (
	"bolg/front-api/global"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"bolg/rpc/article/article_client"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.GetArticleReq) (resp *types.ArticleRespose, err error) {
	article, err := l.svcCtx.Article.GetArticle(l.ctx, &article_client.GetArticleRequest{
		Id: []uint64{req.ID},
	})
	if err != nil {
		return &types.ArticleRespose{
			Code:    200,
			Data:    []*types.ArticleData{},
			Message: global.NULL_DATA,
		}, nil
	}

	if len(article.GetArticleData()) <= 0 {
		return &types.ArticleRespose{
			Code:    200,
			Data:    []*types.ArticleData{},
			Message: global.NULL_DATA,
		}, nil
	}

	resData := article.GetArticleData()[0]
	classifyData := make([]*types.ArticleClassifyData, 0, len(resData.ClassifyData))
	for _, item := range resData.ClassifyData {
		classifyData = append(classifyData, &types.ArticleClassifyData{
			ID:   item.Id,
			Name: item.Name,
			Icon: item.Icon,
		})
	}

	articleData := []*types.ArticleData{
		{
			ID:           resData.Id,
			Title:        resData.Title,
			Content:      resData.Content,
			CreateAt:     resData.CreateAt,
			UpdateAt:     resData.UpdateAt,
			DeleteAt:     resData.DeleteAt,
			Status:       resData.Status,
			Draft:        resData.Draft,
			ClassifyData: classifyData,
		},
	}

	return &types.ArticleRespose{
		Code:    200,
		Data:    articleData,
		Message: global.SUCCESS,
	}, nil
}
