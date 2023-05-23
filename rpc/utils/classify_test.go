package utils

import (
	"bolg/rpc/models"
	"fmt"
	"testing"
)

func TestClassifyConversionType(t *testing.T) {
	articleData := []*models.Article{
		{
			ID:       1,
			Title:    "title",
			Content:  "content",
			CreateAt: 0,
			UpdateAt: 0,
			Status:   0,
		},
	}
	articleData1 := []*models.Article{
		{
			ID:       2,
			Title:    "title2",
			Content:  "content2",
			CreateAt: 1,
			UpdateAt: 2,
			Status:   4,
		},
	}
	modelList := []*models.Classify{
		{
			ID:       1,
			Name:     "zhang",
			CreateAt: 1,
			UpdateAt: 2,
			Icon:     "icon",
			Articles: articleData,
		},
		{
			ID:       2,
			Name:     "hong",
			CreateAt: 1,
			UpdateAt: 1,
			Icon:     "icon2",
			Articles: articleData1,
		},
	}

	classifyData := ClassifyConversionClientType(&modelList)
	fmt.Println(classifyData)

}
