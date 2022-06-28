package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/record"
	"strconv"
)

type RecordHandler struct {
	Logic record.RecordLogic
}

func NewRecordHandler(sqlHandler database.SqlHandler) *RecordHandler {
	return &RecordHandler{
		Logic: record.RecordLogic{
			Repository: &database.RecordRepository{
				Handler: sqlHandler,
			},
		},
	}
}

func (handler *RecordHandler) Get(ctx HttpContext) {
	vars := ctx.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	recordList, err := handler.Logic.GetPersonRecordList(personID)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("指定パーソンのレコード情報が取得できません。"))
		return
	}

	output, _ := json.Marshal(recordList)
	ctx.Response(http.StatusOK, output)
}

func (handler *RecordHandler) Post(ctx HttpContext) {
	vars := ctx.Vars()
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	// TODO: convert value object
	personID, _ := strconv.Atoi(vars["personID"])
	musicID, _ := strconv.Atoi(req["music_id"])
	recordEasy, _ := strconv.Atoi(req["record_easy"])
	recordNormal, _ := strconv.Atoi(req["record_normal"])
	recordHard, _ := strconv.Atoi(req["record_hard"])
	recordExpert, _ := strconv.Atoi(req["record_expert"])
	recordMaster, _ := strconv.Atoi(req["record_master"])

	handler.Logic.Repository.StartTransaction()

	record := model.Record{
		PersonID:     personID,
		MusicID:      musicID,
		RecordEasy:   recordEasy,
		RecordNormal: recordNormal,
		RecordHard:   recordHard,
		RecordExpert: recordExpert,
		RecordMaster: recordMaster,
	}
	recordID, err := handler.Logic.RegistRecord(record)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("レコード情報の登録に失敗しました。"))
		handler.Logic.Repository.Rollback()
		return
	}
	record.RecordID = recordID

	handler.Logic.Repository.Commit()

	output, _ := json.Marshal(record)
	ctx.Response(http.StatusCreated, output)
}

func (handler *RecordHandler) Put(ctx HttpContext) {
	vars := ctx.Vars()
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("レコード情報の登録に失敗しました。"))
		return
	}

	// TODO: convert value object
	personID, _ := strconv.Atoi(vars["personID"])
	musicID, _ := strconv.Atoi(vars["musicID"])
	recordEasy, _ := strconv.Atoi(req["record_easy"])
	recordNormal, _ := strconv.Atoi(req["record_normal"])
	recordHard, _ := strconv.Atoi(req["record_hard"])
	recordExpert, _ := strconv.Atoi(req["record_expert"])
	recordMaster, _ := strconv.Atoi(req["record_master"])

	handler.Logic.Repository.StartTransaction()

	record := model.Record{
		RecordEasy:   recordEasy,
		RecordNormal: recordNormal,
		RecordHard:   recordHard,
		RecordExpert: recordExpert,
		RecordMaster: recordMaster,
	}
	if err := handler.Logic.ModifyRecord(personID, musicID, record); err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("レコード情報の更新に失敗しました。"))
		handler.Logic.Repository.Rollback()
		return
	}

	handler.Logic.Repository.Commit()

	ctx.Response(http.StatusOK, nil)
}
