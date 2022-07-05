package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/logic/inputport"
	"strconv"
)

type RecordHandler struct {
	recordLogic inputport.RecordLogic
}

func NewRecordHandler(r inputport.RecordLogic) *RecordHandler {
	return &RecordHandler{
		recordLogic: r,
	}
}

func (h *RecordHandler) Get(ctx infra.HttpContext) {
	vars := ctx.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	records, err := h.recordLogic.GetByPersonID(personID)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("指定パーソンのレコード情報が取得できません。"))
		return
	}

	output, _ := json.Marshal(records)
	ctx.Response(http.StatusOK, output)
}

func (h *RecordHandler) Post(ctx infra.HttpContext) {
	vars := ctx.Vars()
	var record model.Record
	if err := ctx.Decode(&record); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	// TODO: validation

	// ID check is already finished in middleware check-auth
	personID, _ := strconv.Atoi(vars["personID"])
	record.PersonID = personID

	// recordID, err := h.recordLogic.Store(dum, record)
	// if err != nil {
	// 	ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("レコード情報の登録に失敗しました。"))
	// 	return
	// }
	// record.RecordID = recordID

	output, _ := json.Marshal(record)
	ctx.Response(http.StatusCreated, output)
}

func (h *RecordHandler) Put(ctx infra.HttpContext) {
	vars := ctx.Vars()
	var record model.Record
	if err := ctx.Decode(&record); err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("レコード情報の登録に失敗しました。"))
		return
	}

	// TODO: convert value object
	personID, _ := strconv.Atoi(vars["personID"])
	musicID, _ := strconv.Atoi(vars["musicID"])

	if err := h.recordLogic.Update(personID, musicID, record); err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("レコード情報の更新に失敗しました。"))
		return
	}
	ctx.Response(http.StatusOK, nil)
}
