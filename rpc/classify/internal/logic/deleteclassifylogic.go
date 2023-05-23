package logic

import (
	"bolg/rpc/classify/classify_client"
	"bolg/rpc/classify/internal/svc"
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassifyLogic {
	return &DeleteClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteClassifyLogic) DeleteClassify(in *classify_client.DeleteClassifyRequest) (*classify_client.ClassifyList, error) {
	classify := models.Classify{}
	var classifyList []*models.Classify
	db := common_model.NewDbEnginClient[*models.Classify](l.svcCtx.DbEngin, classify.TableName())

	// 清空关联
	tx := db.DbEngin.Preload("Articles").Find(&classifyList, in.GetId())
	if tx.Error != nil {
		l.Error(tx.Error)
		return nil, tx.Error
	}
	err := db.DbEngin.Model(&classifyList).Association("Articles").Clear()
	if err != nil {
		l.Error(err)
		return nil, err
	}

	// 删除
	tx = db.DbEngin.Delete(&models.Classify{}, in.GetId())
	if tx.Error != nil {
		l.Error(tx.Error)
		return nil, tx.Error
	}
	return &classify_client.ClassifyList{}, nil
}