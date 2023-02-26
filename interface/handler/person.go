package handler

import (
	"context"
	"net/http"
	infraDomain "sekareco_srv/domain/infra"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"strconv"
)

type personHandler struct {
	person inputport.PersonInputport
	valid  inputport.PersonValidator
}

func NewPersonHandler(
	p inputport.PersonInputport,
	v inputport.PersonValidator,
) *personHandler {
	return &personHandler{
		person: p,
		valid:  v,
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
// @Security	Bearer Authentication
// @Router		/persons/{person_id}	[get]
func (h *personHandler) Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	personID, err := infraDomain.GetPersonID(ctx)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	person, err := h.person.GetByID(ctx, personID)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, person)
	return nil
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
func (h *personHandler) Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	var req inputdata.AddPerson
	if err := hc.Decode(&req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.valid.ValidationAdd(ctx, req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	person, err := h.person.Store(ctx, req)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusCreated, person)
	return nil
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
// @Security	Bearer Authentication
// @Router		/persons/{person_id}	[put]
func (h *personHandler) Put(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	var req inputdata.UpdatePerson
	if err := hc.Decode(&req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.valid.ValidationUpdate(ctx, req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusBadRequest}
	}

	if err := h.person.Update(ctx, personID, req); err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, nil)
	return nil
}
