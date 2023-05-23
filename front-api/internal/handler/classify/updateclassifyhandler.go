package classify

import (
	"net/http"

	"bolg/front-api/internal/logic/classify"
	"bolg/front-api/internal/svc"
	"bolg/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateClassifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateClassifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := classify.NewUpdateClassifyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateClassify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
