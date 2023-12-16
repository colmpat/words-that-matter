package db

import (
	"context"
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	// DB is the global database instance
	ProdDB    *DB
	StagingDB *DB
	once      = sync.Once{}
)

type DB struct {
	*bun.DB

	// context for queries
	Ctx context.Context
}

// InitDB initializes the global database instances
func InitDB(prodDSN, stagingDSN string) {
	once.Do(func() {
		ProdDB = MakeDB(prodDSN)
		StagingDB = MakeDB(stagingDSN)
	})
}

// MakeDB makes a new DB instance from a DSN using the pg driver
func MakeDB(dsn string) *DB {
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// make bun db
	bdb := bun.NewDB(pgdb, pgdialect.New())
	// make our db
	return NewDB(bdb)
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
