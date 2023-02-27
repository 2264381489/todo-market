package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"todo-market/todoMarket/api/internal/logic"
	"todo-market/todoMarket/api/internal/svc"
	"todo-market/todoMarket/api/internal/types"
)

func AddTodoLisHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTodoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddTodoLisLogic(r.Context(), svcCtx)
		resp, err := l.AddTodoLis(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
