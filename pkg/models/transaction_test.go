package models

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestWithTransaction(t *testing.T) {

	db, _ := CreateDB("localhost",
		5432,
		"backend",
		"backend",
		"backend",
		"false")

	fn := func(tx *sqlx.Tx) error {
		return nil
	}

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
}
