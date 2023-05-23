package classify

import (
	"net/http"

	"bolg/front-api/internal/logic/classify"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteClassifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteClassifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := classify.NewDeleteClassifyLogic(r.Context(), svcCtx)
		resp, err := l.DeleteClassify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
