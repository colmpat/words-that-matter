package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func validateEnv() {
	req := []string{
		"PROD_DSN",
		"STAGING_DSN",
		"SESSION_SECRET",
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"GOOGLE_REDIRECT_URL",
	}

	godotenv.Load()

	for _, v := range req {
		if _, ok := os.LookupEnv(v); !ok {
			log.Fatalf("Environment variable %s is not set\n", v)
		}
	}
}
