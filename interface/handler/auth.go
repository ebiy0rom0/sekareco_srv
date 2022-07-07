package handler

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
	"strconv"
)

type AuthHandler struct {
	auth inputport.AuthInputport
}

func NewAuthHandler(a inputport.AuthInputport) *AuthHandler {
	return &AuthHandler{
		auth: a,
	}
}

// synonymous with 'sign in'
func (h *AuthHandler) Post(ctx context.Context, hc infra.HttpContext) {
	var req map[string]string
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	personID, err := h.auth.CheckAuth(req["login_id"], req["password"])
	if err != nil {
		hc.Response(http.StatusUnauthorized, hc.MakeError("IDまたはパスワードが間違っています。"))
		return
	}

	token := h.auth.GenerateNewToken()
	// TODO: functionality is provided from infra.AuthMiddleware
	// h.auth.AddToken(personID, token)
}

// synonymous with 'sign out'
func (h *AuthHandler) Delete(ctx context.Context, hc infra.HttpContext) {
	var req map[string]string
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	personID, _ := strconv.Atoi(req["person_id"])
	// TODO: functionality is provided from infra.AuthMiddleware
	// h.auth.RevokeToken(personID)

	hc.Response(http.StatusOK)
}
