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

// @Summary		wip
// @Description get my person and friend person status
// @Tags		persons
// @Accept		json
// @Produce		json
// @param		person_id path	int		true	"Person ID"
// @Success		200
// @Failure		503
// @Router		/prsk/persons/{personID} [get]

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

// @Summary		wip
// @Description wip
// @Tags		accounts
// @Accept		json
// @Produce		json
// @Success		200
// @Failure		503
// @Router		/signup	[post]
func (h *personHandler) Post(ctx context.Context, hc infra.HttpContext) {
	var req inputdata.PostPerson
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

// @Summary		wip
// @Description	person account status change
// @Tags		persons
// @Accept		json
// @Produce		json
// @Success		200
// @Router		/prsk/persons/{personID} [put]
func (h *personHandler) Put(ctx context.Context, hc infra.HttpContext) {
	vars := hc.Vars()
	personID, _ := strconv.Atoi(vars["personID"])

	var req inputdata.PutPerson
	if err := hc.Decode(&req); err != nil {
		hc.Response(http.StatusBadRequest, hc.MakeError(err))
		return
	}

	if err := h.person.Update(ctx, personID, req); err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	hc.Response(http.StatusOK)
}
