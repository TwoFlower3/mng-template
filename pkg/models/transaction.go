package models

import (
	"github.com/jmoiron/sqlx"
)

// TxFn function for query with transaction
type TxFn func(*sqlx.Tx) error

func (db *DB) withTransaction(fn TxFn) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
