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

// Begin starts a transaction.
func (h *txHandler) Begin(ctx context.Context, opt *sql.TxOptions) error {
	tx, err := h.con.BeginTxx(ctx, opt)
	if err != nil {
		return errors.WithStack(err)
	}

	h.tx = tx
	return nil
}

// Execute returns result at execute argument query.
// Prepared statement are supported, so any argument inject to args.
func (h *txHandler) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := h.tx.PreparexContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
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
