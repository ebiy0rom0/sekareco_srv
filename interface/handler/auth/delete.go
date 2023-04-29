package auth

import (
	"context"
	"net/http"
	infraDomain "sekareco_srv/domain/infra"
	"sekareco_srv/interface/infra"
)

// synonymous with 'sign out'
//	@Summary		delete token | delete a stored token
//	@Description	delete a stored token
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Security		Bearer Authentication
//	@Router			/signout	[delete]
func (h *handler) Delete(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	token, err := infraDomain.GetToken(ctx)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusInternalServerError}
	}

	h.authInputport.RevokeToken(token)

	hc.Response(http.StatusOK, nil)
	return nil
}
