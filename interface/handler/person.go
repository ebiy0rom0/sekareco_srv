package handler

import (
	"context"
	"errors"
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

// @Summary		get one | get one person by ID
// @Description	get one person by ID
// @Tags		person
// @Accept		json
// @Produce		json
// @param		person_id	path	int		true	"Want to get person ID"
// @Success		200 {object}	model.Person
// @Failure		503	{object}	infra.HttpError
// @Security	Authentication
// @Router		/persons/{person_id}	[get]
func (h *personHandler) Get(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	person, err := h.person.GetByID(ctx, personID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	hc.Response(http.StatusOK, person)
}

// @Summary		new account | create new person
// @Description	create new person
// @Tags		account
// @Accept		json
// @Produce		json
// @param		data	body	inputdata.AddPerson	true	"add person status"
// @Success		200	{object}	model.Person
// @Failure		400	{object}	infra.HttpError
// @Failure		503	{object}	infra.HttpError
// @Router		/signup	[post]
func (h *personHandler) Post(ctx context.Context, hc infra.HttpContext) {
	var req inputdata.AddPerson
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	if err := req.Valiation(); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	ok, err := h.person.IsDuplicateLoginID(ctx, req.LoginID)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	} else if !ok {
		err = errors.New("loginID is duplicate: " + req.LoginID)
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	person, err := h.person.Store(ctx, req)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	hc.Response(http.StatusCreated, person)
}

// @Summary		update status | update person register status
// @Description	update person register status
// @Tags		person
// @Accept		json
// @Produce		json
// @param		data	body	inputdata.UpdatePerson	true	"update person status"
// @Success		200
// @Failure		400	{object}	infra.HttpError
// @Failure		503	{object}	infra.HttpError
// @Security	Authentication
// @Router		/persons/{person_id}	[put]
func (h *personHandler) Put(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	var req inputdata.UpdatePerson
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	if err := h.person.Update(ctx, personID, req); err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	hc.Response(http.StatusOK, nil)
}
