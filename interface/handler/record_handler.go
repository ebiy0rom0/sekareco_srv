package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sekareco_srv/domain"
	"sekareco_srv/interface/database"
	logic "sekareco_srv/logic/record"
	"strconv"

	"github.com/gorilla/mux"
)

type RecordHandler struct {
	Logic logic.RecordLogic
}

func NewRecordHandler(sqlHandler database.SqlHandler) *RecordHandler {
	return &RecordHandler{
		Logic: logic.RecordLogic{
			Repository: &database.RecordRepository{
				Handler: sqlHandler,
			},
		},
	}
}

func (handler *RecordHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	personId, _ := strconv.Atoi(vars["personId"])

	recordList, err := handler.Logic.GetPersonRecordList(personId)
	if err != nil {
		log.Printf("record list get failed: %s", err)
	}

	output, _ := json.Marshal(recordList)
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (handler *RecordHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	personId, _ := strconv.Atoi(vars["personId"])
	musicId, _ := strconv.Atoi(vars["music_id"])
	recordEasy, _ := strconv.Atoi(vars["record_easy"])
	recordNormal, _ := strconv.Atoi(vars["record_normal"])
	recordHard, _ := strconv.Atoi(vars["record_hard"])
	recordExpert, _ := strconv.Atoi(vars["record_expert"])
	recordMaster, _ := strconv.Atoi(vars["record_master"])

	handler.Logic.Repository.StartTransaction()

	record := domain.Record{
		PersonId:     personId,
		MusicId:      musicId,
		RecordEasy:   recordEasy,
		RecordNormal: recordNormal,
		RecordHard:   recordHard,
		RecordExpert: recordExpert,
		RecordMaster: recordMaster,
	}
	recordId, err := handler.Logic.RegistRecord(record)
	if err != nil {
		log.Printf("record regist failed: %s", err)
		handler.Logic.Repository.Rollback()
		return
	}
	record.RecordId = recordId

	handler.Logic.Repository.Commit()

	output, _ := json.Marshal(record)
	w.WriteHeader(http.StatusCreated)
	w.Write(output)
}

func (handler *RecordHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	personId, _ := strconv.Atoi(vars["personId"])
	musicId, _ := strconv.Atoi(vars["musicId"])
	recordEasy, _ := strconv.Atoi(vars["record_easy"])
	recordNormal, _ := strconv.Atoi(vars["record_normal"])
	recordHard, _ := strconv.Atoi(vars["record_hard"])
	recordExpert, _ := strconv.Atoi(vars["record_expert"])
	recordMaster, _ := strconv.Atoi(vars["record_master"])

	handler.Logic.Repository.StartTransaction()

	record := domain.Record{
		RecordEasy:   recordEasy,
		RecordNormal: recordNormal,
		RecordHard:   recordHard,
		RecordExpert: recordExpert,
		RecordMaster: recordMaster,
	}
	err := handler.Logic.ModifyRecord(personId, musicId, record)
	if err != nil {
		log.Printf("modify record failed: %s", err)
		handler.Logic.Repository.Rollback()
		return
	}

	handler.Logic.Repository.Commit()

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
