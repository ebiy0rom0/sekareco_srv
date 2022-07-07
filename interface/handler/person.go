package handler

import (
	"context"
	"net/http"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
	"strconv"
)

type PersonHandler struct {
	person inputport.PersonInputport
}

func NewPersonHandler(p inputport.PersonInputport) *PersonHandler {
	return &PersonHandler{
		person: p,
	}
}

func (h *PersonHandler) Get(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	person, err := h.person.GetByID(personID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("パーソン情報が取得できません。"))
		return
	}

	hc.Response(http.StatusOK, person)
}

// synonymous with 'sign out'
func (h *PersonHandler) Post(ctx context.Context, hc infra.HttpContext) {
	var req model.PostPerson
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	// TODO: request parameter validation
	// if err != nil {
	// 	hc.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータが不正です。"))
	// 	return
	// }

	ok, err := h.person.IsDuplicateLoginID(req.LoginID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("重複チェックの検証に失敗しました。"))
		return
	} else if !ok {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("ログインIDは既に存在しています。"))
		return
	}

	person, err := h.person.Store(req)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("パーソンの登録に失敗しました。"))
		return
	}

	hc.Response(http.StatusCreated, person)
}

// TODO: Implement
func (h *PersonHandler) Put(ctx context.Context, hc infra.HttpContext) {

	vars := hc.Vars()

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "put"
	// output, _ := json.Marshal(vars)

	hc.Response(http.StatusOK, vars)
}
