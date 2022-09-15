package sql

import (
	"context"
	"database/sql"
	"sekareco_srv/interface/infra"
)

// A txHandler is database handler wrapper supports the transaction.
type txHandler struct {
	con *sql.DB
	tx  *sql.Tx
}

// NewTxHandler returns txHandler pointer.
func NewTxHandler (con *sql.DB) *txHandler {
	return &txHandler{con: con}
}

// Begin starts a transaction.
func (h *txHandler) Begin(ctx context.Context, opt *sql.TxOptions) error {
	tx, err := h.con.BeginTx(ctx, opt)
	if err != nil {
		return err
	}

	h.tx = tx
	return nil
}

// Execute returns result at execute argument query.
// Prepared statement are supported, so any argument inject to args.
func (h *txHandler) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := h.tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Commit commits the transaction.
func (h *txHandler) Commit() error {
	return h.tx.Commit()
}

// Rollback aborts a transaction.
func (h *txHandler) Rollback() error {
	return h.tx.Rollback()
}

var _ infra.TxHandler = (*txHandler)(nil)