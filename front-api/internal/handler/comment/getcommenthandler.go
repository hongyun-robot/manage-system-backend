package comment

//func GetCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req types.GetCommentReq
//		if err := httpx.Parse(r, &req); err != nil {
//			httpx.ErrorCtx(r.Context(), w, err)
//			return
//		}
//
//		l := comment.NewGetCommentLogic(r.Context(), svcCtx)
//		resp, err := l.GetComment(&req)
//		if err != nil {
//			httpx.ErrorCtx(r.Context(), w, err)
//		} else {
//			httpx.OkJsonCtx(r.Context(), w, resp)
//		}
//	}
//}
