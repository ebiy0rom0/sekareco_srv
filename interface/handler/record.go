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

	// recordID, err := h.recordLogic.Store(dum, record)
	// if err != nil {
	// 	hc.Response(http.StatusServiceUnavailable, hc.MakeError("レコード情報の登録に失敗しました。"))
	// 	return
	// }
	// record.RecordID = recordID

	output, _ := json.Marshal(record)
	hc.Response(http.StatusCreated, output)
}

func (h *recordHandler) Put(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	var record inputdata.PutRecord
	if err := hc.Decode(&record); err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
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
