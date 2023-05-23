package article

import (
	"bolg/front-api/global"
	"bolg/rpc/article/article_client"
	"context"
	"fmt"

	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleReq) (resp *types.ArticleRespose, err error) {
	// todo: add your logic here and delete this line
	article, err := l.svcCtx.Article.AddArticle(l.ctx, &article_client.AddArticleRequest{
		Title:       req.Title,
		Content:     req.Content,
		Status:      req.Status,
		Draft:       req.Draft,
		ClassifyIds: req.ClassifyIds,
	})
	if err != nil {
		fmt.Println("front-api add err", err)
		return nil, err
	}

	if len(article.GetArticleData()) <= 0 {
		return &types.ArticleRespose{
			Code:    200,
			Data:    []*types.ArticleData{},
			Message: global.SUCCESS,
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