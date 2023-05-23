package utils

import (
	"bolg/rpc/base/base_client"
	"fmt"
	"testing"
)

func TestMenuRpcToGorm(t *testing.T) {
	rpcData := []*base_client.MenuData{
		{
			Title:    "title",
			Icon:     "icon",
			Url:      "url",
			Disabled: false,
			Hide:     false,
			Children: []*base_client.MenuData{
				{
					Id:       11,
					Title:    "title",
					Icon:     "icon",
					Url:      "url",
					Disabled: false,
					Hide:     false,
					Children: nil,
				},
				{
					Id:       21,
					Title:    "title",
					Icon:     "icon",
					Url:      "url",
					Disabled: false,
					Hide:     false,
					Children: nil,
				},
			},
		},
		{
			Id:       2,
			Title:    "title",
			Icon:     "icon",
			Url:      "url",
			Disabled: false,
			Hide:     false,
			Children: []*base_client.MenuData{
				{
					Id:       21,
					Title:    "title",
					Icon:     "icon",
					Url:      "url",
					Disabled: false,
					Hide:     false,
					Children: nil,
				},
				{
					Id:       22,
					Title:    "title",
					Icon:     "icon",
					Url:      "url",
					Disabled: false,
					Hide:     false,
					Children: nil,
				},
			},
		},
	}
	data, err := MenuRpcToGorm(rpcData)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, datum := range data {
		fmt.Println(i, datum)
	}

	fmt.Println(data[0].Children[0])
}