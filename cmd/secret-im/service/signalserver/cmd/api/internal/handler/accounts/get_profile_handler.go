package handler

import (
	"net/http"

	"secret-im/service/signalserver/cmd/api/internal/logic/accounts"
	"secret-im/service/signalserver/cmd/api/internal/svc"
	"secret-im/service/signalserver/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetProfileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetProfileLogic(r.Context(), ctx)
		resp, err := l.GetProfile(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}