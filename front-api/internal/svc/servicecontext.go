package svc

import (
	"bolg/front-api/internal/config"
	"bolg/rpc/article/article"
	"bolg/rpc/base/base"
	"bolg/rpc/classify/classify"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Article  article.Article
	Classify classify.Classify
	Base     base.Base
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Article:  article.NewArticle(zrpc.MustNewClient(c.ArticleRpc)),
		Classify: classify.NewClassify(zrpc.MustNewClient(c.ClassifyRpc)),
		Base:     base.NewBase(zrpc.MustNewClient(c.BaseRpc)),
	}
}