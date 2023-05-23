package logic

import (
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"bolg/rpc/utils"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"bolg/rpc/classify/classify_client"
	"bolg/rpc/classify/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassifyLogic {
	return &UpdateClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateClassifyLogic) UpdateClassify(in *classify_client.UpdateClassifyRequest) (*classify_client.ClassifyList, error) {
	classifyData := models.Classify{
		ID:       in.Id,
		Name:     in.Name,
		Icon:     in.Icon,
		Articles: nil,
	}
	articleList := make([]*models.Article, 0, len(in.ArticleIds))

	db := common_model.NewDbEnginClient[*models.Classify](l.svcCtx.DbEngin, classifyData.TableName())

	// 调用获取方法，获取该 ID 分类所有数据
	getClassify := NewGetClassifyLogic(l.ctx, l.svcCtx)
	getClassifyRequest := classify_client.GetClassifyRequest{Id: []uint64{in.GetId()}}
	classifyList, err := getClassify.GetClassify(&getClassifyRequest)
	if err != nil {
		return nil, err
	}

	// 将 rpc 数据类型转为 model 数据类型
	classifyListData := classifyList.GetClassify()
	classifyModelData := utils.ClassifyConversionType(&classifyListData)[0]
	// 清空所有关联关系 -- 再重新创建关联关系
	err = db.DbEngin.Model(classifyModelData).Association("Articles").Clear()
	if err != nil {
		fmt.Println("清空关联失败", err)
		l.Errorf("清空关联关系失败 err=%v", err)
		return nil, err
	}

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

	db.DbEngin.Save(&classifyData)

	classifyResponseData := utils.ClassifyConversionClientType(&[]*models.Classify{
		&classifyData,
	})

	return classifyResponseData, nil
}