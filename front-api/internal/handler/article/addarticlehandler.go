package article

import (
	"net/http"

	"bolg/front-api/internal/logic/article"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewAddArticleLogic(r.Context(), svcCtx)
		resp, err := l.AddArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
