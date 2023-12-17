package main

import (
	"os"

	"github.com/colmpat/words-that-matter/internal/auth"
	"github.com/colmpat/words-that-matter/pkg/db"
	"github.com/gin-gonic/gin"
)

type Link struct {
	URL  string
	Text string
}

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
	return NewAdminHandler(&Config{
		r.Group("/admin"),
	})
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

	makeAuthHandler(r)
	makeAdminHandler(r)

	r.Run()
}
