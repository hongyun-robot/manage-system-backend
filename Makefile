all: run_etcd run_rpc run_api

run_etcd:
	etcd

run_api:
	go run ./front-api/front.go -f ./etc/front-api.yaml

run_rpc: run_rpc_classify run_rpc_article
run_rpc_base:
	go run ./rpc/base/base.go -f ./rpc/base/etc/base.yaml
run_rpc_classify:
	go run ./rpc/classify/classify.go -f ./rpc/classify/etc/classify.yaml
run_rpc_article:
	go run ./rpc/article/article.go -f ./rpc/article/etc/article.yaml