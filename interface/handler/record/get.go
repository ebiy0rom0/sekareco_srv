package record

import (
	"context"
	"fmt"
	"net/http"
	"sekareco_srv/interface/infra"
	"strconv"
)

//	@Summary		get list | get all records data by person
//	@Description	get all records data by person
//	@Tags			record
//	@Accept			json
//	@Produce		json
//	@param			person_id	path		int	true	"Want to get personID"
//	@Success		200			{object}	[]model.Record
//	@Failure		503			{object}	infra.HttpError
//	@Security		Bearer Authentication
//	@Router			/records/{person_id}	[get]
func (h *handler) Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	vars := hc.Vars()
	fmt.Printf("%+v\n", vars)
	personID, _ := strconv.Atoi(vars["person_id"])

	records, err := h.recordInputport.GetByPersonID(ctx, personID)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, records)
	return nil
}
