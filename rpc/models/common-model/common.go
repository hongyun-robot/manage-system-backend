package common_model

import (
	"bolg/rpc/models"
	"gorm.io/gorm"
)

type ModelsBase interface {
	TableName() string
}

type BaseTypeInterface interface {
	ModelsBase
	*models.Article | *models.Classify
}

type BaseType[T BaseTypeInterface] []T

type DbEnginClient[T BaseTypeInterface] interface {
	Insert(data *T) (tx *gorm.DB, err error)
}

type BaseDbEnginClient[T BaseTypeInterface] struct {
	DbEngin   *gorm.DB
	TableName string
}

func NewDbEnginClient[T BaseTypeInterface](dbEngin *gorm.DB, tableName string) *BaseDbEnginClient[T] {
	return &BaseDbEnginClient[T]{
		DbEngin:   dbEngin,
		TableName: tableName,
	}
}

func (receiver *BaseDbEnginClient[T]) Insert(data T) (tx *gorm.DB, err error) {
	result := receiver.DbEngin.Table(receiver.TableName).Create(data)
	return result, nil
}
func (receiver *BaseDbEnginClient[T]) Find(data T) (tx *gorm.DB) {
	tx = receiver.DbEngin.Table(receiver.TableName).First(data)
	return tx
}
