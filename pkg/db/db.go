package db

import (
	"context"

	"github.com/uptrace/bun"
)

type DB struct {
	*bun.DB

	// context for queries
	Ctx context.Context
}

// NewDB returns a new DB instance. Creates a background context for queries.
func NewDB(db *bun.DB) *DB {
	return &DB{
		DB:  db,
		Ctx: context.Background(),
	}
}

// Transact executes the given function in a transaction. If the function returns an error, the transaction is rolled back. If the function panics, the transaction is rolled back and the panic is re-thrown.
func (db *DB) Transact(txFunc func(tx *bun.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = txFunc(&tx)
	return err
}
