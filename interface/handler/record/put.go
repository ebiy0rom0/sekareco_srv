package record

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"strconv"
)

//	@Summary		update status | update record clear status
//	@Description	update record clear status
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@param			person_id	path	int						true	"Want to update personID"
//	@param			music_id	path	int						true	"Want to update musicID"
//	@param			data		body	inputdata.UpdateRecord	true	"update Record"
//	@Success		200
//	@Failure		400	{object}	infra.HttpError
//	@Failure		503	{object}	infra.HttpError
//	@Security		Bearer Authentication
//	@Router			/records/{person_id}/{music_id}	[put]
func (h *handler) Put(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	vars := hc.Vars()
	var record inputdata.UpdateRecord
	if err := hc.Decode(&record); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	// TODO: convert value object
	personID, _ := strconv.Atoi(vars["personID"])
	musicID, _ := strconv.Atoi(vars["musicID"])

	if err := h.recordInputport.Update(ctx, personID, musicID, record); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, nil)
	return nil
}
