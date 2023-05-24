package article

import (
	"bolg/front-api/global"
	"bolg/front-api/internal/utils"
	"bolg/rpc/article/article_client"
	"context"
	"fmt"
	"net/http"

	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByPagingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByPagingLogic {
	return &GetArticleByPagingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByPagingLogic) GetArticleByPaging(req *types.GetArticleByPagingReq) (resp *types.ArticleRespose, err error) {
	fmt.Println("draft", req.Draft)
	articleList, err := l.svcCtx.Article.GetArticleByPaging(l.ctx, &article_client.GetArticleByPagingRequest{
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		Draft:    req.Draft,
	})
	fmt.Println("qqqqqqqqq", err)
	if err != nil {
		return &types.ArticleRespose{
			Code:    http.StatusOK,
			Data:    []*types.ArticleData{},
			Message: global.ERROR,
		}, nil
	}

	articleData, err := utils.ArticleRpcToApi(articleList.GetArticleData())
	if err != nil {
		return &types.ArticleRespose{
			Code:    http.StatusOK,
			Data:    []*types.ArticleData{},
			Message: global.ERROR,
		}, nil
	}
	return &types.ArticleRespose{
		Code:    http.StatusOK,
		Data:    articleData,
		Message: global.SUCCESS,
	}, nil
}
