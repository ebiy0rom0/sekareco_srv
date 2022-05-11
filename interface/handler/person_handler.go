package handler

import (
	"fmt"
	"log"
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

func (handler *PersonHandler) Get(c HttpContext) {
	vars := c.Vars()
	personId, _ := strconv.Atoi(vars["personId"])

	person, err := handler.logic.GetPersonById(personId)
	if err != nil {
		log.Printf("%s", err)
	}

	c.Response(http.StatusOK, person)
}

func (handler *PersonHandler) Post(c HttpContext) {
	var p map[string]string
	c.Decode(&p)
	fmt.Println(p)
	vo, err := dto.NewRequestRegistPerson(p["login_id"], p["password"], p["person_name"])
	if err != nil {
		log.Printf("bad request: %s", err)
	}

	ok, err := handler.logic.CheckDuplicateLoginId(vo.GetLoginId())
	if err != nil {
		log.Printf("%s", err)
	} else if !ok {
		log.Printf("duplicate login id: %s", vo.GetLoginId())
	}

	handler.logic.Repository.StartTransaction()

	code := handler.logic.GenerateFriendCode(vo.GetLoginId())
	person := model.Person{
		PersonName: vo.GetPersonName(),
		FriendCode: code,
	}
	personId, err := handler.logic.RegistPerson(person)
	if err != nil {
		log.Printf("person regist failed: %s", err)
		handler.logic.Repository.Rollback()
		return
	}
	person.PersonId = personId

	hash, err := handler.logic.GeneratePasswordHash(vo.GetPassword())
	if err != nil {
		log.Printf("password hash generate failed: %s", err)
	}
	login := model.Login{
		LoginId:      vo.GetLoginId(),
		PersonId:     personId,
		PasswordHash: hash,
	}

	err = handler.logic.RegistLogin(login)
	if err != nil {
		log.Printf("login regist failed: %s", err)
		handler.logic.Repository.Rollback()
		return
	}

	handler.logic.Repository.Commit()

	c.Response(http.StatusCreated, person)
}

// TODO: Implement
func (handler *PersonHandler) Put(c HttpContext) {

	vars := c.Vars()

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "put"
	// output, _ := json.Marshal(vars)

	c.Response(http.StatusOK, vars)
}
