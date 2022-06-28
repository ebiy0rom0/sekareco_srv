package handler

import (
	"net/http"
	"sekareco_srv/infra/middleware"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/auth"
)

type AuthHandler struct {
	logic auth.AuthLogic
}

func NewAuthHandler(sqlHandler database.SqlHandler) *AuthHandler {
	return &AuthHandler{
		logic: auth.AuthLogic{
			Repository: &database.AuthRepository{
				Handler: sqlHandler,
			},
		},
	}
}

// synonymous with 'sign in'
func (h *AuthHandler) Post(ctx HttpContext) {
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	personID, err := h.logic.CheckAuth(req["login_id"], req["password"])
	if err != nil {
		ctx.Response(http.StatusUnauthorized, ctx.MakeError("IDまたはパスワードが間違っています。"))
		return
	}

	token := ""
	middleware.Auth.AddTokens(personID, token)
}

// synonymous with 'sign out'
func (handler *AuthHandler) Delete(ctx HttpContext) {
}
