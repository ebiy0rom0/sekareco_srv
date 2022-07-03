package handler

import (
	"net/http"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/auth"
	"strconv"
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

	token := h.logic.GenerateNewToken()
	h.logic.AddToken(personID, token)
}

// synonymous with 'sign out'
func (h *AuthHandler) Delete(ctx HttpContext) {
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	personID, _ := strconv.ParseInt(req["person_id"], 10, 8)
	h.logic.RevokeToken(int(personID))

	ctx.Response(http.StatusOK)
}
