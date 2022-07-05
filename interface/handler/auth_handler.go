package handler

import (
	"net/http"
	"sekareco_srv/domain/model"
	"strconv"
)

type AuthHandler struct {
	authLogic model.AuthLogic
}

func NewAuthHandler(a model.AuthLogic) *AuthHandler {
	return &AuthHandler{
		authLogic: a,
	}
}

// synonymous with 'sign in'
func (h *AuthHandler) Post(ctx HttpContext) {
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	personID, err := h.authLogic.CheckAuth(req["login_id"], req["password"])
	if err != nil {
		ctx.Response(http.StatusUnauthorized, ctx.MakeError("IDまたはパスワードが間違っています。"))
		return
	}

	token := h.authLogic.GenerateNewToken()
	h.authLogic.AddToken(personID, token)
}

// synonymous with 'sign out'
func (h *AuthHandler) Delete(ctx HttpContext) {
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	personID, _ := strconv.ParseInt(req["person_id"], 10, 8)
	h.authLogic.RevokeToken(int(personID))

	ctx.Response(http.StatusOK)
}
