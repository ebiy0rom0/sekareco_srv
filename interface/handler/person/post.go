package person

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
)

//	@Summary		new account | create new person
//	@Description	create new person
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@param			data	body		inputdata.AddPerson	true	"add person status"
//	@Success		200		{object}	model.Person
//	@Failure		400		{object}	infra.HttpError
//	@Failure		503		{object}	infra.HttpError
//	@Router			/signup	[post]
func (h *handler) Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	var req inputdata.AddPerson
	if err := hc.Decode(&req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.personValidator.ValidationAdd(ctx, req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	person, err := h.personInputport.Store(ctx, req)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusCreated, person)
	return nil
}
