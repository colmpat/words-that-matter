package main

import (
	"database/sql"
	"os"

	"github.com/colmpat/words-that-matter/internal/auth"
	"github.com/colmpat/words-that-matter/pkg/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

type (
	Link struct {
		URL  string
		Text string
	}
)

func makeAuthHandler(r *gin.Engine) *AuthHandler {
	g := r.Group("/auth")
	providers := []auth.Provider{
		auth.NewGoogleProvider(),
		// can add more providers here
	}

	// build the auth handler
	return NewAuthHandler(&Config{
		g,
	}, providers...)
}

func makeAdminHandler(r *gin.Engine) *AdminHandler {
	g := r.Group("/admin")
	g.Use(auth.AdminAuthTo("/auth/google/login", "/"))

	return NewAdminHandler(&Config{
		g,
	})
}

func initSessionStore(DB *sql.DB) postgres.Store {
	store, err := postgres.NewStore(DB, []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	return store
}

func init() {
	validateEnv()
	db.InitDB(
		os.Getenv("PROD_DSN"),
		os.Getenv("STAGING_DSN"),
	)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	store := initSessionStore(db.ProdDB.DB.DB)
	r.Use(sessions.Sessions("session", store))
	makeAuthHandler(r)
	makeAdminHandler(r)

	r.Run()
}
