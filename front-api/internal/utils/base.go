package utils

import (
	"bolg/front-api/internal/types"
	"bolg/rpc/base/base_client"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
)

// MenuApiToRpc 接口参数转为 rpc 类型
func MenuApiToRpc(data []*types.MenuData) (menuData []*base_client.MenuData, err error) {
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

// MenuRpcToApi 将 rpc 返回数据类型转为 api 响应类型
func MenuRpcToApi(data []*base_client.MenuData) (menuData []*types.MenuData, err error) {
	menuData = make([]*types.MenuData, 0, len(data))
	for _, datum := range data {
		menu := types.MenuData{}
		err := convertor.CopyProperties(&menu, datum)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		menuData = append(menuData, &menu)
	}
	return menuData, nil
}