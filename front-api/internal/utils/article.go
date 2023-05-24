package utils

import (
	"bolg/front-api/internal/types"
	"bolg/rpc/article/article_client"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
)

// ArticleRpcToApi 将 rpc 返回数据类型转为 api 响应类型
func ArticleRpcToApi(data []*article_client.ArticleData) (articleData []*types.ArticleData, err error) {
	articleData = make([]*types.ArticleData, 0, len(data))
	for _, datum := range data {
		article := types.ArticleData{}
		err := convertor.CopyProperties(&article, datum)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		articleData = append(articleData, &article)
	}
	return articleData, nil
}
