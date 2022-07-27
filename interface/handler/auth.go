package handler

import (
	"context"
	"fmt"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"strconv"
)

type authHandler struct {
	auth inputport.AuthInputport
}

func NewAuthHandler(a inputport.AuthInputport) *authHandler {
	return &authHandler{
		auth: a,
	}
}

// synonymous with 'sign in'
// @Summary		add token | generate and stored token
// @Description	generate and stored token
// @Tags		account
// @Accept		json
// @Produce		json
// @param		login_id	body	string	true	"LoginID"
// @param		password	body	string	true	"password"
// @Success		200	{string}	string	"generate new token"
// @Failure		400	{object}	infra.HttpError
// @Failure		401	{object}	infra.HttpError
// @Router		/signin	[post]
func (h *authHandler) Post(ctx context.Context, hc infra.HttpContext) {
	var req inputdata.PostAuth
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	if err := req.Validation(); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
	}

	personID, err := h.auth.CheckAuth(ctx, req.LoginID, req.Password)
	if err != nil {
		hc.Response(http.StatusUnauthorized, hc.MakeError(err))
		return
	}

	token := h.auth.GenerateNewToken()
	// TODO: functionality is provided from infra.AuthMiddleware
	// h.auth.AddToken(personID, token)
	fmt.Printf("add token: personID->%d, token->%s", personID, token)
}

// synonymous with 'sign out'
// @Summary		delete token | delete a stored token
// @Description	delete a stored token
// @Tags		account
// @Accept		json
// @Produce		json
// @param		person_id		body	string	true	"Want to delete token personID"
// @Success		200
// @Failure		400	{object}	infra.HttpError
// @Security	Authentication
// @Router		/signout	[delete]
func (h *authHandler) Delete(ctx context.Context, hc infra.HttpContext) {
	var req inputdata.DeleteAuth
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	personID, _ := strconv.Atoi(req.PersonID)
	// TODO: functionality is provided from infra.AuthMiddleware
	// h.auth.RevokeToken(personID)
	fmt.Printf("revoke token: personID->%d", personID)

	hc.Response(http.StatusOK)
}
