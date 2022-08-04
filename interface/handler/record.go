package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"strconv"
)

type recordHandler struct {
	record inputport.RecordInputport
}

func NewRecordHandler(r inputport.RecordInputport) *recordHandler {
	return &recordHandler{
		record: r,
	}
}

// @Summary		get list | get all records data by person
// @Description	get all records data by person
// @Tags		record
// @Accept		json
// @Produce		json
// @param		person_id	path	int		true	"Want to get personID"
// @Success		200	{object}	[]model.Record
// @Failure		503	{object}	infra.HttpError
// @Security	Bearer Authentication
// @Router		/records/{person_id}	[get]
func (h *recordHandler) Get(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	fmt.Printf("%+v\n", vars)
	personID, _ := strconv.Atoi(vars["person_id"])

	records, err := h.record.GetByPersonID(ctx, personID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	output, _ := json.Marshal(records)
	hc.Response(http.StatusOK, output)
}

// @Summary		new record | create new record
// @Description	create new record
// @Tags		record
// @Accept		json
// @Produce		json
// @param		person_id	path	int						true	"Want to add personID"
// @param		data		body	inputdata.AddRecord		true	"store Record"
// @Success		201	{object}	model.Record
// @Failure		400	{object}	infra.HttpError
// @Failure		503	{object}	infra.HttpError
// @Security	Bearer Authentication
// @Router		/records/{person_id}	[post]
func (h *recordHandler) Post(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	var addRecord inputdata.AddRecord
	if err := hc.Decode(&addRecord); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	// TODO: validation

	// ID check is already finished in middleware check-auth
	personID, _ := strconv.Atoi(vars["person_id"])

	newRecord, err := h.record.Store(ctx, personID, addRecord)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	output, _ := json.Marshal(newRecord)
	hc.Response(http.StatusCreated, output)
}

// @Summary		update status | update record clear status
// @Description	update record clear status
// @Tags		record
// @Accept		json
// @Produce		json
// @param		person_id	path	int						true	"Want to update personID"
// @param		music_id	path	int						true	"Want to update musicID"
// @param		data		body	inputdata.UpdateRecord	true	"update Record"
// @Success		200
// @Failure		400	{object}	infra.HttpError
// @Failure		503	{object}	infra.HttpError
// @Security	Bearer Authentication
// @Router		/records/{person_id}/{music_id}	[put]
func (h *recordHandler) Put(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	var record inputdata.UpdateRecord
	if err := hc.Decode(&record); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	// TODO: convert value object
	personID, _ := strconv.Atoi(vars["personID"])
	musicID, _ := strconv.Atoi(vars["musicID"])

	if err := h.record.Update(ctx, personID, musicID, record); err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}
	hc.Response(http.StatusOK, nil)
}
