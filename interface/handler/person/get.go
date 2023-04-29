package person

import (
	"context"
	"net/http"
	infraDomain "sekareco_srv/domain/infra"
	"sekareco_srv/interface/infra"
)

//	@Summary		get one | get one person by ID
//	@Description	get one person by ID
//	@Tags			person
//	@Accept			json
//	@Produce		json
//	@param			person_id	path		int	true	"Want to get person ID"
//	@Success		200			{object}	model.Person
//	@Failure		503			{object}	infra.HttpError
//	@Security		Bearer Authentication
//	@Router			/persons/{person_id}	[get]
func (h *handler) Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	personID, err := infraDomain.GetPersonID(ctx)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	person, err := h.personInputport.GetByID(ctx, personID)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, person)
	return nil
}
