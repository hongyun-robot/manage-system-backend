package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySQL struct {
		DataSourceName    string
		DefaultStringSize uint
	}
}