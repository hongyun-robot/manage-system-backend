package logic

import (
	"bolg/rpc/article/article_client"
	"bolg/rpc/article/internal/svc"
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddArticleLogic) AddArticle(in *article_client.AddArticleRequest) (*article_client.ArticleList, error) {
	classifyLen := len(in.ClassifyIds)
	classifyList := make([]*models.Classify, 0, classifyLen)
	articleData := &models.Article{
		Title:     in.Title,
		Content:   in.Content,
		Status:    in.Status,
		Draft:     in.Draft,
		Classifys: classifyList,
	}
	db := common_model.NewDbEnginClient[*models.Article](l.svcCtx.DbEngin, articleData.TableName())

	for _, id := range in.ClassifyIds {
		var classify models.Classify
		result := db.DbEngin.Table(classify.TableName()).First(&classify, id)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			fmt.Println("暂无数据")
			fmt.Println("err", result.Error)
		} else {
			classifyList = append(classifyList, &classify)
		}
	}
	articleData.Classifys = classifyList
	result := db.DbEngin.Create(articleData)
	if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		fmt.Println("保存失败")
		fmt.Println("err", result.Error)
		return nil, result.Error
	}

	classifyClientData := make([]*article_client.ClassifyData, 0, classifyLen)
	for _, classify := range classifyList {
		classifyClientData = append(classifyClientData, &article_client.ClassifyData{
			Id:       classify.ID,
			Name:     classify.Name,
			Icon:     classify.Icon,
			CreateAt: classify.CreateAt,
			UpdateAt: classify.UpdateAt,
		})
	}

	var articleResponseData []*article_client.ArticleData

	articleResponseData = append(articleResponseData, &article_client.ArticleData{
		Id:           articleData.ID,
		Title:        articleData.Title,
		Content:      articleData.Content,
		Status:       articleData.Status,
		Draft:        articleData.Draft,
		CreateAt:     articleData.CreateAt,
		UpdateAt:     articleData.UpdateAt,
		ClassifyData: classifyClientData,
	})

	return &article_client.ArticleList{
		ArticleData: articleResponseData,
	}, nil
}
