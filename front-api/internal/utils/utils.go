package utils

import (
	"bolg/front-api/internal/types"
	"bolg/rpc/article/article_client"
	"bolg/rpc/classify/classify_client"
	"golang.org/x/exp/constraints"
)

// ClassifyConversionType classify 类型转换 classify_client 转为 api
func ClassifyConversionType(classifyList *[]*classify_client.ClassifyData) []*types.ClassifyData {
	data := make([]*types.ClassifyData, 0, len(*classifyList))

	for _, classifyClientData := range *classifyList {
		articleList := make([]*types.ClassifyArticleData, 0, len(classifyClientData.ArticleData))
		for _, article := range classifyClientData.ArticleData {
			articleData := &types.ClassifyArticleData{
				ID:       article.Id,
				Title:    article.Title,
				Content:  article.Content,
				UpdateAt: article.UpdateAt,
				Status:   article.Status,
				Draft:    article.Draft,
			}

			articleList = append(articleList, articleData)
		}
		classifyData := &types.ClassifyData{
			ID:          classifyClientData.Id,
			Name:        classifyClientData.Name,
			CreateAt:    classifyClientData.CreateAt,
			UpdateAt:    classifyClientData.UpdateAt,
			Icon:        classifyClientData.Icon,
			ArticleData: articleList,
		}

		data = append(data, classifyData)
	}

	return data
}

// ArticleConversionType article 类型转换 article_client 转为 api
func ArticleConversionType(articleList *[]*article_client.ArticleData) []*types.ArticleData {
	data := make([]*types.ArticleData, 0, len(*articleList))

	for _, articleClientData := range *articleList {
		classifyList := make([]*types.ArticleClassifyData, 0, len(articleClientData.ClassifyData))
		for _, classify := range articleClientData.ClassifyData {
			classifyData := &types.ArticleClassifyData{
				ID:   classify.Id,
				Name: classify.Name,
				Icon: classify.Icon,
			}

			classifyList = append(classifyList, classifyData)
		}
		articleData := &types.ArticleData{
			ID:           articleClientData.Id,
			Title:        articleClientData.Title,
			Content:      articleClientData.Content,
			CreateAt:     articleClientData.CreateAt,
			UpdateAt:     articleClientData.UpdateAt,
			DeleteAt:     articleClientData.DeleteAt,
			Status:       articleClientData.Status,
			Draft:        articleClientData.Draft,
			ClassifyData: classifyList,
		}

		data = append(data, articleData)
	}

	return data
}

// RemoveRepeatedElement 去重
func RemoveRepeatedElement[T constraints.Ordered](s []T) []T {
	result := make([]T, 0)
	m := make(map[T]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}
