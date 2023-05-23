package svc

import (
	"bolg/rpc/article/internal/config"
	"bolg/rpc/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.MySQL.DataSourceName
	defaultStringSize := c.MySQL.DefaultStringSize
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         defaultStringSize,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库链接报错", err)
		return nil
	}

	err = db.AutoMigrate(&models.Article{}, &models.Classify{})
	if err != nil {
		fmt.Println("建表失败-article", err)
		panic("AutoMigrate err")
		return nil
	}
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}