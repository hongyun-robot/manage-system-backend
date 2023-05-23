package utils

import (
	"bolg/rpc/models"
	"fmt"
	"testing"
)

func TestArticleConversionClientType(t *testing.T) {
	classifyData1 := []*models.Classify{
		{
			ID:       1,
			Name:     "classifyData1",
			CreateAt: 123,
			UpdateAt: 456,
			Icon:     "icon2",
		},
	}
	classifyData2 := []*models.Classify{
		{
			ID:       1,
			Name:     "classifyData1",
			CreateAt: 123,
			UpdateAt: 456,
			Icon:     "icon2",
		},
	}

	modelList := []*models.Article{
		{
			ID:        2,
			Title:     "title1",
			Content:   "content",
			CreateAt:  123,
			UpdateAt:  321,
			Status:    1,
			Classifys: classifyData1,
		},
		{
			ID:        3,
			Title:     "title1",
			Content:   "content",
			CreateAt:  123,
			UpdateAt:  321,
			Status:    1,
			Classifys: classifyData2,
		},
	}

	articleList := ArticleConversionClientType(&modelList)
	fmt.Println(articleList)
}