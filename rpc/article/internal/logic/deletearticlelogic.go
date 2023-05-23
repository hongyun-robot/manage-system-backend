package logic

import (
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"context"
	"errors"
	"fmt"

	"bolg/rpc/article/article_client"
	"bolg/rpc/article/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleLogic) DeleteArticle(in *article_client.DeleteArticleRequest) (*article_client.ArticleList, error) {
	var article models.Article
	var articleList []*models.Article

	db := common_model.NewDbEnginClient[*models.Article](l.svcCtx.DbEngin, article.TableName())

	tx := db.DbEngin.Unscoped().Find(&articleList, in.GetId())
	fmt.Println(articleList[0].DeleteAt.Valid)
	if tx.Error != nil {
		l.Error(tx.Error)
		return nil, tx.Error
	}

	in.Id = []uint64{}

	for _, articleData := range articleList {
		if !articleData.DeleteAt.Valid {
			in.Id = append(in.Id, articleData.ID)
		}
	}

	if len(in.GetId()) <= 0 {
		return nil, errors.New("NULL DATA")
	}

	err := db.DbEngin.Delete(&article, in.GetId()).Error
	fmt.Println(err)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	return &article_client.ArticleList{}, nil
}