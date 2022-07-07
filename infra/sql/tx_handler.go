package sql

import "database/sql"

type txHandler struct {
	tx *sql.Tx
}

func newTxHandler(tx *sql.Tx) *txHandler {
	return &txHandler{
		tx: tx,
	}
}

func (h *txHandler) Commit() error {
	return h.tx.Commit()
}

func (h *txHandler) Rollback() error {
	return h.tx.Rollback()
}

func (h *txHandler) Execute(query string, args ...interface{}) (sql.Result, error) {
	h.tx.ExecContext()
	// stmt, err := h.tx.PrepareContext(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	// res, err := stmt.Exec(args...)
	// if err != nil {
	// 	return res, err
	// }
	// return res, nil
}
