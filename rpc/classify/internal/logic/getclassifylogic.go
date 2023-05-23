package logic

import (
	"bolg/rpc/classify/classify_client"
	"bolg/rpc/classify/internal/svc"
	"bolg/rpc/models"
	common_model "bolg/rpc/models/common-model"
	"bolg/rpc/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifyLogic {
	return &GetClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassifyLogic) GetClassify(in *classify_client.GetClassifyRequest) (*classify_client.ClassifyList, error) {
	var classify models.Classify
	db := common_model.NewDbEnginClient[*models.Classify](l.svcCtx.DbEngin, classify.TableName())

	var modelsClassifyList []*models.Classify

	tx := db.DbEngin.Preload("Articles").Find(&modelsClassifyList, in.GetId())

	if tx.Error != nil {
		l.Error(tx.Error)
		return nil, tx.Error
	}

	// 将 model 类型转为响应类型
	classifyResponseData := utils.ClassifyConversionClientType(&modelsClassifyList)

	return classifyResponseData, nil
}