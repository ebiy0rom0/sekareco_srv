package sql

import (
	"context"
	"database/sql"
)

type txHandler struct {
	con *sql.DB
	tx  *sql.Tx
}

func NewTxHandler (con *sql.DB) *txHandler {
	return &txHandler{con: con}
}

func (h *txHandler) Begin(ctx context.Context, opt *sql.TxOptions) error {
	tx, err := h.con.BeginTx(ctx, opt)
	if err != nil {
		return err
	}

	h.tx = tx
	return nil
}

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

func (h *txHandler) Commit() error {
	return h.tx.Commit()
}

func (h *txHandler) Rollback() error {
	return h.tx.Rollback()
}
