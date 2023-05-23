package logic

import (
	"bolg/rpc/classify/classify_client"
	"bolg/rpc/classify/internal/svc"
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"bolg/rpc/utils"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type AddClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassifyLogic {
	return &AddClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddClassifyLogic) AddClassify(in *classify_client.AddClassifyRequest) (*classify_client.ClassifyList, error) {
	modelsClassifyLst := make([]*models.Classify, 0, 1)
	articleList := make([]*models.Article, 0, len(in.ArticleIds))
	classifyData := models.Classify{
		Name:     in.Name,
		Icon:     in.Icon,
		Articles: articleList,
	}

	db := common_model.NewDbEnginClient[*models.Classify](l.svcCtx.DbEngin, classifyData.TableName())

	// 查询分类数据，用以关联文章
	for _, id := range in.ArticleIds {
		var article models.Article
		result := db.DbEngin.Table(article.TableName()).First(&article, id)
		if errors.As(result.Error, &gorm.ErrRecordNotFound) {
			fmt.Println("暂无数据")
			fmt.Println("db.DbEngin.First(articleList)", result.Error)
		} else {
			articleList = append(articleList, &article)
		}
	}

	classifyData.Articles = articleList

	// 创建数据
	result := db.DbEngin.Create(&classifyData)
	if errors.As(result.Error, &gorm.ErrRecordNotFound) {
		fmt.Println("添加失败")
		return nil, result.Error
	}

	modelsClassifyLst = append(modelsClassifyLst, &classifyData)

	// 将 model 类型转为响应类型
	classifyResponseData := utils.ClassifyConversionClientType(&modelsClassifyLst)
	return classifyResponseData, nil
}
