package auth

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
)

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
func (h *handler) Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	var req inputdata.PostAuth
	if err := hc.Decode(&req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.authValidator.ValidationPost(req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	personID, err := h.authInputport.CheckAuth(ctx, req.LoginID, req.Password)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusUnauthorized}
	}

	token := h.authInputport.AddToken(personID)

	hc.Response(http.StatusOK, token)
	return nil
}
