package main

import (
	"log"
	"os"

	"github.com/colmpat/words-that-matter/pkg/db"
	"github.com/colmpat/words-that-matter/pkg/models"
	"github.com/joho/godotenv"
)

// parseArgs parses the command line arguments and returns whether to migrate, prod, staging, or all
func parseArgs() (prod bool, staging bool, all bool) {
	for _, arg := range os.Args {
		switch arg {
		case "--prod", "-p":
			prod = true
		case "--staging", "-s":
			staging = true
		case "--all", "-a":
			all = true
			return
		}
	}

	return
}

// registerModels registers the models with the provided database
func registerModels(DB *db.DB, condition bool, name string) {
	if !condition {
		return
	}
	log.Println("Registering models for database:", name)

	models := []interface{}{
		&models.Media{},
		&models.Deck{},
		&models.Frequency{},
		&models.Term{},
	}

	err := DB.ResetModel(DB.Ctx, models...)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	prod, staging, all := parseArgs()

	godotenv.Load()
	db.InitDB(os.Getenv("PROD_DSN"), os.Getenv("STAGING_DSN"))

	registerModels(db.ProdDB, all || prod, "prod")
	registerModels(db.StagingDB, all || staging, "staging")
}
