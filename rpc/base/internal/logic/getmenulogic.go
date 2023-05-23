package logic

import (
	"bolg/rpc/models"
	"bolg/rpc/utils"
	"context"
	"fmt"
	"gorm.io/gorm"

	"bolg/rpc/base/base_client"
	"bolg/rpc/base/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMenu rpc AddArticle(AddArticleRequest) returns(ArticleList);
func (l *GetMenuLogic) GetMenu(in *base_client.GetMenuRequest) (*base_client.MenuListResponse, error) {
	fmt.Println(in)
	db := l.svcCtx.DbEngin.Session(&gorm.Session{})
	var data []*models.Menu
	data, err := GetAllMenu(db)
	if err != nil {
		return nil, err
	}

	responseData, err := utils.MenuGormToRpc(data)
	if err != nil {
		fmt.Println("err", err)
		l.Errorf("GetMenu -> MenuGormToRpc;err=", err)
		return nil, err
	}

	return &base_client.MenuListResponse{
		MenuData: responseData,
	}, nil
}

func GetAllMenu(db *gorm.DB) (menuList []*models.Menu, err error) {
	if err = db.Where("parent_id is null").Order("`order`").Find(&menuList).Error; err != nil {
		return nil, err
	}
	for _, menu := range menuList {
		children, err := GetRecursionMenu(db, menu.Id)
		if err != nil {
			return nil, err
		}
		menu.Children = children
	}
	return menuList, nil
}

func GetRecursionMenu(db *gorm.DB, id uint64) (menuList []*models.Menu, err error) {
	if err = db.Where("parent_id = ?", id).Order("`order`").Find(&menuList).Error; err != nil {
		return
	}
	for _, menu := range menuList {
		menu.Children, err = GetRecursionMenu(db, menu.Id)
		if err != nil {
			return
		}
	}
	return
}
