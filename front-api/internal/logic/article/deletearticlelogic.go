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

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (resp *types.ArticleRespose, err error) {
	article, err := l.svcCtx.Article.DeleteArticle(l.ctx, &article_client.DeleteArticleRequest{Id: []uint64{req.ID}})
	if err != nil {
		return &types.ArticleRespose{
			Code:    200,
			Data:    nil,
			Message: global.NULL_DATA,
		}, nil
	}

	fmt.Println(article)

	return &types.ArticleRespose{
		Code:    200,
		Data:    nil,
		Message: global.SUCCESS,
	}, nil
}
