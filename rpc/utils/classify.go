package utils

import (
	"bolg/rpc/classify/classify_client"
	"bolg/rpc/models"
)

// ClassifyConversionClientType classify 类型转换 model 转为 classify_client
func ClassifyConversionClientType(modelList *[]*models.Classify) *classify_client.ClassifyList {
	data := classify_client.ClassifyList{Classify: nil}
	classifyData := make([]*classify_client.ClassifyData, 0, len(*modelList))
	for _, classify := range *modelList {
		articleList := make([]*classify_client.ArticleData, 0, len(classify.Articles))
		for _, article := range classify.Articles {
			articleList = append(articleList, &classify_client.ArticleData{
				Id:       article.ID,
				Title:    article.Title,
				Content:  article.Content,
				Status:   article.Status,
				Draft:    article.Draft,
				CreateAt: article.CreateAt,
				UpdateAt: article.UpdateAt,
			})
		}

		classifyData = append(classifyData, &classify_client.ClassifyData{
			Id:          classify.ID,
			Name:        classify.Name,
			Icon:        classify.Icon,
			CreateAt:    classify.CreateAt,
			UpdateAt:    classify.UpdateAt,
			ArticleData: articleList,
		})
	}

	data.Classify = classifyData

	return &data
}

// ClassifyConversionType classify 类型转换 classify_client 转为 model
func ClassifyConversionType(classifyList *[]*classify_client.ClassifyData) []*models.Classify {
	data := make([]*models.Classify, 0, len(*classifyList))

	for _, classifyClientData := range *classifyList {
		articleList := make([]*models.Article, 0, len(classifyClientData.ArticleData))
		for _, article := range classifyClientData.ArticleData {
			articleData := &models.Article{
				ID:       article.Id,
				Title:    article.Title,
				Content:  article.Content,
				CreateAt: article.CreateAt,
				UpdateAt: article.UpdateAt,
				Status:   article.Status,
			}

			articleList = append(articleList, articleData)
		}
		classifyData := &models.Classify{
			ID:       classifyClientData.Id,
			Name:     classifyClientData.Name,
			CreateAt: classifyClientData.CreateAt,
			UpdateAt: classifyClientData.UpdateAt,
			Icon:     classifyClientData.Icon,
			Articles: articleList,
		}

		data = append(data, classifyData)
	}

	return data
}