package handler

import (
	"context"
	"net/http"
	infraDomain "sekareco_srv/domain/infra"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
)

type authHandler struct {
	auth  inputport.AuthInputport
	valid inputport.AuthValidator
}

type AuthHandle interface {
	Post(context.Context, infra.HttpContext) *infra.HttpError
	Delete(context.Context, infra.HttpContext) *infra.HttpError
}

func NewAuthHandler(
	a inputport.AuthInputport,
	v inputport.AuthValidator,
) *authHandler {
	return &authHandler{
		auth:  a,
		valid: v,
	}
}

// synonymous with 'sign in'
//	@Summary		add token | generate and stored token
//	@Description	generate and stored token
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@param			data	body		inputdata.PostAuth	true	"password"
//	@Success		200		{string}	string				"generate new token"
//	@Failure		400		{object}	infra.HttpError
//	@Failure		401		{object}	infra.HttpError
//	@Router			/signin	[post]
func (h *authHandler) Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	var req inputdata.PostAuth
	if err := hc.Decode(&req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.valid.ValidationPost(req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	personID, err := h.auth.CheckAuth(ctx, req.LoginID, req.Password)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusUnauthorized}
	}

	token := h.auth.AddToken(personID)

	hc.Response(http.StatusOK, token)
	return nil
}

// synonymous with 'sign out'
//	@Summary		delete token | delete a stored token
//	@Description	delete a stored token
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Security		Bearer Authentication
//	@Router			/signout	[delete]
func (h *authHandler) Delete(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	token, err := infraDomain.GetToken(ctx)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	h.auth.RevokeToken(token)

	hc.Response(http.StatusOK, nil)
	return nil
}

var _ AuthHandle = (*authHandler)(nil)
