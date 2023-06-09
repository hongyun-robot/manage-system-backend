// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	article "bolg/front-api/internal/handler/article"
	basemenu "bolg/front-api/internal/handler/base/menu"
	classify "bolg/front-api/internal/handler/classify"
	"bolg/front-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: basemenu.GetMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: basemenu.AddMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: basemenu.UpdateMenuHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: basemenu.DeleteMenuHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/base/menu"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: article.AddArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: article.GetArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getarticle",
				Handler: article.GetArticleByPagingHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: article.UpdateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: article.DeleteArticleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/article"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: classify.AddClassifyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: classify.GetClassifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: classify.UpdateClassifyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: classify.DeleteClassifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/classify"),
	)
}
