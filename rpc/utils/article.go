package utils

import (
	"bolg/rpc/article/article_client"
	"bolg/rpc/models"
)

// ArticleConversionClientType article 类型转换 model 转为 article_client
func ArticleConversionClientType(modelList *[]*models.Article) *article_client.ArticleList {
	data := article_client.ArticleList{ArticleData: nil}
	articleData := make([]*article_client.ArticleData, 0, len(*modelList))
	for _, article := range *modelList {
		classifyData := make([]*article_client.ClassifyData, 0, len(article.Classifys))
		for _, classify := range article.Classifys {
			classifyData = append(classifyData, &article_client.ClassifyData{
				Id:       classify.ID,
				Name:     classify.Name,
				Icon:     classify.Icon,
				CreateAt: classify.CreateAt,
				UpdateAt: classify.UpdateAt,
			})
		}

		articleData = append(articleData, &article_client.ArticleData{
			Id:           article.ID,
			Title:        article.Title,
			Content:      article.Content,
			Status:       article.Status,
			Draft:        article.Draft,
			CreateAt:     article.CreateAt,
			UpdateAt:     article.UpdateAt,
			ClassifyData: classifyData,
		})
	}

	data.ArticleData = articleData

	return &data
}

// ArticleConversionType article 类型转换 article_client 转为 model
func ArticleConversionType(articleList *[]*article_client.ArticleData) []*models.Article {
	data := make([]*models.Article, 0, len(*articleList))

	for _, articleClientData := range *articleList {
		classifyList := make([]*models.Classify, 0, len(articleClientData.ClassifyData))
		for _, classify := range articleClientData.ClassifyData {
			classifyData := &models.Classify{
				ID:       classify.Id,
				Name:     classify.Name,
				CreateAt: classify.CreateAt,
				UpdateAt: classify.UpdateAt,
				Icon:     classify.Icon,
			}

			classifyList = append(classifyList, classifyData)
		}
		articleData := &models.Article{
			ID:        articleClientData.Id,
			Title:     articleClientData.Title,
			Content:   articleClientData.Content,
			CreateAt:  articleClientData.CreateAt,
			UpdateAt:  articleClientData.UpdateAt,
			Status:    articleClientData.Status,
			Classifys: classifyList,
		}

		data = append(data, articleData)
	}

	return data
}
