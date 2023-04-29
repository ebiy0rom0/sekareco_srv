package record

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"strconv"
)

//	@Summary		new record | create new record
//	@Description	create new record
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@param			person_id	path		int					true	"Want to add personID"
//	@param			data		body		inputdata.AddRecord	true	"store Record"
//	@Success		201			{object}	model.Record
//	@Failure		400			{object}	infra.HttpError
//	@Failure		503			{object}	infra.HttpError
//	@Security		Bearer Authentication
//	@Router			/records/{person_id}	[post]
func (h *handler) Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	vars := hc.Vars()
	var addRecord inputdata.AddRecord
	if err := hc.Decode(&addRecord); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	// TODO: validation

	// ID check is already finished in middleware check-auth
	personID, _ := strconv.Atoi(vars["person_id"])

	newRecord, err := h.recordInputport.Store(ctx, personID, addRecord)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusCreated, newRecord)
	return nil
}
