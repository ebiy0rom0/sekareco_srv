package handler

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"strconv"
)

type personHandler struct {
	person inputport.PersonInputport
}

func NewPersonHandler(p inputport.PersonInputport) *personHandler {
	return &personHandler{
		person: p,
	}
}

func (h *personHandler) Get(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	person, err := h.person.GetByID(ctx, personID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("パーソン情報が取得できません。"))
		return
	}

	hc.Response(http.StatusOK, person)
}

// synonymous with 'sign out'
func (h *personHandler) Post(ctx context.Context, hc infra.HttpContext) {
	var req inputdata.PostPerson
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	if err := req.Valiation(); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err.Error()))
		return
	}

	ok, err := h.person.IsDuplicateLoginID(ctx, req.LoginID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("重複チェックの検証に失敗しました。"))
		return
	} else if !ok {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("ログインIDは既に存在しています。"))
		return
	}

	person, err := h.person.Store(ctx, req)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("パーソンの登録に失敗しました。"))
		return
	}

	hc.Response(http.StatusCreated, person)
}

func (h *personHandler) Put(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	var req inputdata.PutPerson
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError("failed to decode request parameter"))
		return
	}

	if err := h.person.Update(ctx, personID, req); err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err.Error()))
		return
	}

	hc.Response(http.StatusOK)
}
