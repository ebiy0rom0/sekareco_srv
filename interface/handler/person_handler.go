package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sekareco_srv/domain"
	"sekareco_srv/interface/database"
	logic "sekareco_srv/logic/person"
	"strconv"

	"github.com/gorilla/mux"
)

type PersonHandler struct {
	logic logic.PersonLogic
}

func NewPersonHandler(sqlHandler database.SqlHandler) *PersonHandler {
	return &PersonHandler{
		logic: logic.PersonLogic{
			Repository: &database.PersonRepository{
				Handler: sqlHandler,
			},
		},
	}
}

func (handler *PersonHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	personId, _ := strconv.Atoi(vars["personId"])

	person, err := handler.logic.GetPersonById(personId)
	if err != nil {
		log.Printf("%s", err)
	}

	output, _ := json.Marshal(person)
	w.Write(output)
}

func (handler *PersonHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	loginId := vars["login_id"]

	// TODO: no rows return error
	ok, err := handler.logic.CheckDuplicateLoginId(loginId)
	if err != nil {
		log.Printf("%s", err)
	} else if !ok {
		log.Printf("duplicate login id: %s", loginId)
	}

	code := handler.logic.GenerateFriendCode(loginId)

	person := domain.Person{
		PersonName: vars["person_name"],
		FriendCode: code,
	}
	personId, err := handler.logic.RegistPerson(person)
	if err != nil {
		log.Printf("person regist failed: %s", err)
	}
	person.PersonId = personId

	password := vars["password"]
	hash, err := handler.logic.GeneratePasswordHash(password)
	if err != nil {
		log.Printf("password hash generate failed: %s", err)
	}

	login := domain.Login{
		LoginId:      loginId,
		PersonId:     personId,
		PasswordHash: hash,
	}
	err = handler.logic.RegistLogin(login)
	if err != nil {
		log.Printf("login regist failed: %s", err)
	}

	output, _ := json.Marshal(person)
	w.Write(output)
}

// TODO: Implement
func (handler *PersonHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	// @debug
	vars["uri"] = "person"
	vars["methods"] = "put"
	output, _ := json.Marshal(vars)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
