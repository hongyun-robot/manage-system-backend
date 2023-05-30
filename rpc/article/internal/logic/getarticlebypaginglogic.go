package logic

import (
	"bolg/rpc/article/article_client"
	"bolg/rpc/article/internal/svc"
	"bolg/rpc/models"
	"bolg/rpc/utils"
	"context"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByPagingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleByPagingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByPagingLogic {
	return &GetArticleByPagingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleByPagingLogic) GetArticleByPaging(in *article_client.GetArticleByPagingRequest) (*article_client.ArticleList, error) {
	//fmt.Println(in.)
	db := l.svcCtx.DbEngin.Session(&gorm.Session{})
	var data []*models.Article
	if in.Draft != nil {
		db = db.Where("draft = ?", in.GetDraft())
	}
	if err := db.Preload("Classifys").Limit(int(in.PageSize)).Offset(int((in.GetPageNum() - 1) * in.GetPageSize())).Find(&data).Error; err != nil {
		l.Error(err)
		return nil, err
	}

	articleResponseData := utils.ArticleConversionClientType(&data)

	return articleResponseData, nil
}
