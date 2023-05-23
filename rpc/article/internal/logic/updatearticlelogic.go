package logic

import (
	"bolg/rpc/article/article_client"
	"bolg/rpc/article/internal/svc"
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"bolg/rpc/utils"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateArticleLogic) UpdateArticle(in *article_client.UpdateArticleRequest) (*article_client.ArticleList, error) {
	articleData := models.Article{
		ID:        in.Id,
		Title:     in.Title,
		Content:   in.Content,
		Status:    in.Status,
		Draft:     in.Draft,
		Classifys: nil,
	}
	classifyList := make([]*models.Classify, 0, len(in.ClassifyIds))

	db := common_model.NewDbEnginClient[*models.Article](l.svcCtx.DbEngin, articleData.TableName())

	// 调用获取方法，获取该 ID 文章所有数据
	getArticle := NewGetArticleLogic(l.ctx, l.svcCtx)
	getArticleRequest := article_client.GetArticleRequest{Id: []uint64{in.GetId()}}
	articleList, err := getArticle.GetArticle(&getArticleRequest)
	if err != nil {
		return nil, err
	}

	// 将 rpc 数据类型转为 model 数据类型
	articleListData := articleList.GetArticleData()
	articleModelData := utils.ArticleConversionType(&articleListData)[0]
	// 清空所有关联关系 -- 再重新创建关联关系
	err = db.DbEngin.Model(articleModelData).Association("Classifys").Clear()
	if err != nil {
		fmt.Println("清空关联失败", err)
		l.Errorf("清空关联关系失败 err=%v", err)
		return nil, err
	}

	// 查询文章数据，用以关联分类
	db.DbEngin.Find(&classifyList, in.ClassifyIds)
	articleData.Classifys = classifyList

	db.DbEngin.Save(&articleData)

	articleResponseData := utils.ArticleConversionClientType(&[]*models.Article{
		&articleData,
	})

	return articleResponseData, nil
}
