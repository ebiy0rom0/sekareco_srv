package handler

import (
	"net/http"
	"sekareco_srv/domain/model"
	"strconv"
)

type PersonHandler struct {
	personLogic model.PersonLogic
}

func NewPersonHandler(l model.PersonLogic) *PersonHandler {
	return &PersonHandler{
		personLogic: l,
	}
}

func (h *PersonHandler) Get(ctx HttpContext) {
	vars := ctx.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	person, err := h.personLogic.GetByID(personID)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("パーソン情報が取得できません。"))
		return
	}

	ctx.Response(http.StatusOK, person)
}

func (h *PersonHandler) Post(ctx HttpContext) {
	var req model.PostPerson
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	// TODO: request parameter validation
	// if err != nil {
	// 	ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータが不正です。"))
	// 	return
	// }

	ok, err := h.personLogic.IsDuplicateLoginID(req.LoginID)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("重複チェックの検証に失敗しました。"))
		return
	} else if !ok {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("ログインIDは既に存在しています。"))
		return
	}

	person, err := h.personLogic.Store(req)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("パーソンの登録に失敗しました。"))
		return
	}

	ctx.Response(http.StatusCreated, person)
}

// TODO: Implement
func (h *PersonHandler) Put(ctx HttpContext) {

	vars := ctx.Vars()

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "put"
	// output, _ := json.Marshal(vars)

	ctx.Response(http.StatusOK, vars)
}
