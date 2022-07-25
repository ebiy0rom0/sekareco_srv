package handler

import (
	"context"
	"encoding/json"
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
// @Tags		records
// @Accept		json
// @Produce		json
// @param		person_id		query	int		true	"Want to get personID"
// @Success		200	{object}	[]model.Record
// @Failure		503	{object}	infra.HttpError
// @Security	Authentication
// @Router		/prsk/records/{person_id}	[get]
func (h *recordHandler) Get(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

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
// @Tags		records
// @Accept		json
// @Produce		json
// @param		person_id		query	int		true	"Want to add personID"
// @param		music_id		body	int		true	"store target musicID"
// @param		record_easy		body	int		false	"easy's clear status"		enums(0,1,2,3)
// @param		record_normal	body	int		false	"normal's clear status"		enums(0,1,2,3)
// @param		record_hard		body	int		false	"hard's clear status"		enums(0,1,2,3)
// @param		record_expert	body	int		false	"expert's clear status"		enums(0,1,2,3)
// @param		record_master	body	int		false	"master's clear status"		enums(0,1,2,3)
// @Success		200	{object}	model.Record
// @Failure		400	{object}	infra.HttpError
// @Failure		503	{object}	infra.HttpError
// @Security	Authentication
// @Router		/prsk/records/{person_id}	[post]
func (h *recordHandler) Post(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	var record inputdata.PostRecord
	if err := hc.Decode(&record); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	// TODO: validation

	// ID check is already finished in middleware check-auth
	personID, _ := strconv.Atoi(vars["personID"])
	record.PersonID = personID

	recordID, err := h.record.Store(ctx, record)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}
	record.RecordID = recordID

	output, _ := json.Marshal(record)
	hc.Response(http.StatusCreated, output)
}

// @Summary		update status | update record clear status
// @Description	update record clear status
// @Tags		records
// @Accept		json
// @Produce		json
// @param		person_id		query	int		true	"Want to update personID"
// @param		music_id		query	int		true	"Want to update musicID"
// @param		record_easy		body	int		false	"easy's clear status"		enums(0,1,2,3)
// @param		record_normal	body	int		false	"normal's clear status"		enums(0,1,2,3)
// @param		record_hard		body	int		false	"hard's clear status"		enums(0,1,2,3)
// @param		record_expert	body	int		false	"expert's clear status"		enums(0,1,2,3)
// @param		record_master	body	int		false	"master's clear status"		enums(0,1,2,3)
// @Success		200	{object}	[]model.Record
// @Failure		400	{object}	infra.HttpError
// @Failure		503	{object}	infra.HttpError
// @Security	Authentication
// @Router		/prsk/records/{person_id}/{music_id}	[put]
func (h *recordHandler) Put(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	var record inputdata.PutRecord
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
