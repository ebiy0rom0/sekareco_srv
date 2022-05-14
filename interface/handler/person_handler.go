package handler

import (
	"net/http"
	"sekareco_srv/domain/dto"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/person"
	"strconv"
)

type PersonHandler struct {
	logic person.PersonLogic
}

func NewPersonHandler(sqlHandler database.SqlHandler) *PersonHandler {
	return &PersonHandler{
		logic: person.PersonLogic{
			Repository: &database.PersonRepository{
				Handler: sqlHandler,
			},
		},
	}
}

func (handler *PersonHandler) Get(ctx HttpContext) {
	vars := ctx.Vars()
	personId, _ := strconv.Atoi(vars["personId"])

	person, err := handler.logic.GetPersonById(personId)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("パーソン情報が取得できません。"))
		return
	}

	ctx.Response(http.StatusOK, person)
}

func (handler *PersonHandler) Post(ctx HttpContext) {
	var req map[string]string
	if err := ctx.Decode(&req); err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータの取得に失敗しました。"))
		return
	}

	vo, err := dto.NewRequestRegistPerson(req["login_id"], req["password"], req["person_name"])
	if err != nil {
		ctx.Response(http.StatusBadRequest, ctx.MakeError("リクエストパラメータが不正です。"))
		return
	}

	ok, err := handler.logic.CheckDuplicateLoginId(vo.GetLoginId())
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("重複チェックの検証に失敗しました。"))
		return
	} else if !ok {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("ログインIDは既に存在しています。"))
		return
	}

	handler.logic.Repository.StartTransaction()

	code, _ := handler.logic.GenerateFriendCode(vo.GetLoginId())
	person := model.Person{
		PersonName: vo.GetPersonName(),
		FriendCode: code,
	}
	personId, err := handler.logic.RegistPerson(person)
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("パーソン情報の登録に失敗しました。"))
		handler.logic.Repository.Rollback()
		return
	}
	person.PersonId = personId

	hash, err := handler.logic.GeneratePasswordHash(vo.GetPassword())
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("パスワードハッシュの生成に失敗しました。"))
		return
	}
	login := model.Login{
		LoginId:      vo.GetLoginId(),
		PersonId:     personId,
		PasswordHash: hash,
	}

	if err := handler.logic.RegistLogin(login); err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("ログイン情報の登録に失敗しました。"))
		handler.logic.Repository.Rollback()
		return
	}

	handler.logic.Repository.Commit()

	ctx.Response(http.StatusCreated, person)
}

// TODO: Implement
func (handler *PersonHandler) Put(ctx HttpContext) {

	vars := ctx.Vars()

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "put"
	// output, _ := json.Marshal(vars)

	ctx.Response(http.StatusOK, vars)
}
