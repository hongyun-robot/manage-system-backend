package logic

import (
	"bolg/rpc/article/article_client"
	"bolg/rpc/article/internal/svc"
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"bolg/rpc/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleLogic) GetArticle(in *article_client.GetArticleRequest) (*article_client.ArticleList, error) {
	db := common_model.NewDbEnginClient[*models.Article](l.svcCtx.DbEngin, "article")

	var modelsArticleList []*models.Article

	tx := db.DbEngin.Preload("Classifys").Find(&modelsArticleList, in.GetId())

	if tx.Error != nil {
		l.Error(tx.Error)
		return nil, tx.Error
	}

	articleResponseData := utils.ArticleConversionClientType(&modelsArticleList)

	return articleResponseData, nil
}
