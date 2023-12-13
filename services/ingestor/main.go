package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/colmpat/words-that-matter/pkg/db"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func validateEnv() {
	req := []string{
		"DSN",
	}

	godotenv.Load()

	for _, v := range req {
		if _, ok := os.LookupEnv(v); !ok {
			log.Fatalf("Environment variable %s is not set\n", v)
		}
	}
}

func main() {
	validateEnv()

	dsn := os.Getenv("DSN")
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// make bun db
	bdb := bun.NewDB(pgdb, pgdialect.New())
	// make our db
	db := db.NewDB(bdb)

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	ib := NewIngestorBuilder()
	ib.DB(db)
	ib.Interval(10 * time.Second)

	ing, err := ib.Build()
	if err != nil {
		log.Fatalf("Error building ingestor: %s\n", err)
	}
	ing.Start()

	<-make(chan struct{})
}
