package logic

import (
	"bolg/rpc/classify/classify_client"
	"bolg/rpc/classify/internal/svc"
	"bolg/rpc/models"
	"bolg/rpc/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllClassifyLogic {
	return &GetAllClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllClassifyLogic) GetAllClassify() (*classify_client.ClassifyList, error) {
	var modelsClassifyList []*models.Classify
	tx := l.svcCtx.DbEngin.Find(&modelsClassifyList)
	if tx.Error != nil {
		l.Error(tx.Error)
		return nil, tx.Error
	}

	// 将 model 类型转为响应类型
	classifyResponseData := utils.ClassifyConversionClientType(&modelsClassifyList)

	return classifyResponseData, nil
}
