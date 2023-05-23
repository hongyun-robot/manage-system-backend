package logic

import (
	"bolg/rpc/base/base_client"
	"bolg/rpc/base/internal/svc"
	"bolg/rpc/models"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMenuLogic) DeleteMenu(in *base_client.DeleteMenuRequest) (*base_client.MenuListResponse, error) {
	for _, id := range in.GetId() {
		if err := DeleteMenuTree(l.svcCtx.DbEngin, id); err != nil {
			l.Errorf("DeleteMenu -> DeleteMenuTree; err=%v", err)
			return nil, err
		}
	}
	return &base_client.MenuListResponse{}, nil
}

// DeleteMenuTree 递归删除
func DeleteMenuTree(db *gorm.DB, id uint64) error {
	// 查询数据
	var menu models.Menu
	if err := db.First(&menu, id).Error; err != nil {
		fmt.Println("==============")
		fmt.Println(errors.As(err, &gorm.ErrRecordNotFound))
		return err
	}

	// 查询所有子数据
	var children []models.Menu

	if err := db.Where("parent_id = ?", menu.Id).Find(&children).Error; err != nil {
		return err
	}

	// 递归删除子数据
	for _, child := range children {
		if err := DeleteMenuTree(db, child.Id); err != nil {
			return err
		}
	}

	// 删除本条数据
	if err := db.Delete(&menu).Error; err != nil {
		return err
	}

	return nil
}