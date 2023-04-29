package person

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"strconv"
)

//	@Summary		update status | update person register status
//	@Description	update person register status
//	@Tags			person
//	@Accept			json
//	@Produce		json
//	@param			data	body	inputdata.UpdatePerson	true	"update person status"
//	@Success		200
//	@Failure		400	{object}	infra.HttpError
//	@Failure		503	{object}	infra.HttpError
//	@Security		Bearer Authentication
//	@Router			/persons/{person_id}	[put]
func (h *handler) Put(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	var req inputdata.UpdatePerson
	if err := hc.Decode(&req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.personValidator.ValidationUpdate(ctx, req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.personInputport.Update(ctx, personID, req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, nil)
	return nil
}
