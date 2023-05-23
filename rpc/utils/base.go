package utils

import (
	"bolg/rpc/base/base_client"
	"bolg/rpc/models"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
)

func MenuRpcToGorm(data []*base_client.MenuData) (menuData []*models.Menu, err error) {
	menuData = make([]*models.Menu, 0, len(data))
	for _, datum := range data {
		menu := models.Menu{}
		err := convertor.CopyProperties(&menu, datum)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		menuData = append(menuData, &menu)
	}
	return menuData, nil
}
func MenuGormToRpc(data []*models.Menu) (menuData []*base_client.MenuData, err error) {
	menuData = make([]*base_client.MenuData, 0, len(data))
	for _, datum := range data {
		menu := base_client.MenuData{}
		err := convertor.CopyProperties(&menu, datum)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		menuData = append(menuData, &menu)
	}
	return menuData, nil
}
