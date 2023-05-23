all: gen_rpc_base gen_rpc_article gen_rpc_classify

gen_rpc_base:
	goctl rpc protoc ./base/base.proto --go_out=..\..\rpc\base --go-grpc_out=..\..\rpc\base  --zrpc_out=..\..\rpc\base
gen_rpc_article:
	goctl rpc protoc ./article/article.proto --go_out=..\..\rpc\article --go-grpc_out=..\..\rpc\article  --zrpc_out=..\..\rpc\article
gen_rpc_classify:
	goctl rpc protoc ./classify/classify.proto --go_out=..\..\rpc\classify --go-grpc_out=..\..\rpc\classify  --zrpc_out=..\..\rpc\classify