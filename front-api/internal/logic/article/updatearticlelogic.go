package article

import (
	"bolg/front-api/global"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"bolg/front-api/internal/utils"
	"bolg/rpc/article/article_client"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (resp *types.ArticleRespose, err error) {
	article, err := l.svcCtx.Article.UpdateArticle(l.ctx, &article_client.UpdateArticleRequest{
		Id:          req.Id,
		Title:       req.Title,
		Content:     req.Content,
		Status:      req.Status,
		Draft:       req.Draft,
		ClassifyIds: utils.RemoveRepeatedElement[uint64](req.ClassifyIds),
	})
	if err != nil {
		l.Errorf("rpc 查询失败 UpdateArticle id=", req.Id)
		return &types.ArticleRespose{
			Code:    200,
			Data:    nil,
			Message: global.NULL_DATA,
		}, nil
	}

	articleData := utils.ArticleConversionType(&article.ArticleData)

	return &types.ArticleRespose{
		Code:    200,
		Data:    articleData,
		Message: global.SUCCESS,
	}, nil
}
