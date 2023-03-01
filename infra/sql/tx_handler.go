package sql

import (
	"context"
	"database/sql"
	"sekareco_srv/interface/infra"

	"github.com/ebiy0rom0/errors"
	"github.com/jmoiron/sqlx"
)

// A txHandler is database handler wrapper supports the transaction.
type txHandler struct {
	con *sqlx.DB
	tx  *sqlx.Tx
}

// NewTxHandler returns txHandler pointer.
func NewTxHandler(con *sqlx.DB) *txHandler {
	return &txHandler{con: con}
}

// BeginTxx start a transaction.
func (h *txHandler) BeginTxx(ctx context.Context, opt *sql.TxOptions) error {
	tx, err := h.con.BeginTxx(ctx, opt)
	if err != nil {
		return errors.WithStack(err)
	}

	h.tx = tx
	return nil
}

// ExecNamedContext returns result at execute argument query.
// Any named placeholder parameters are replaced with fields from arg.
// To exec update query, use 'UpdateNamedContext()'.
func (h *txHandler) ExecNamedContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	stmt, err := h.tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, arg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

// UpdateNamedContext returns result at execute argument query.
// Named placeholder in update field parameters are replaced with fields from param.
// Use placeholders other than named parameters in the where clause.
// To exec Insert and Delete query, use 'ExecNamedContext()'.
func (h *txHandler) UpdateNamedContext(ctx context.Context, query string, param interface{}, args ...interface{}) (sql.Result, error) {
	query, list, err := sqlx.Named(query, param)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	list = append(list, args...)

	stmt, err := h.con.PreparexContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, list...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

// Commit commits the transaction.
func (h *txHandler) Commit() error {
	if err := h.tx.Commit(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Rollback aborts a transaction.
func (h *txHandler) Rollback() error {
	if err := h.tx.Rollback(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

var _ infra.TxHandler = (*txHandler)(nil)
