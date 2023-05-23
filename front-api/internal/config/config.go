package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DataSourceName string
	ArticleRpc     zrpc.RpcClientConf
	ClassifyRpc    zrpc.RpcClientConf
	BaseRpc        zrpc.RpcClientConf
}
