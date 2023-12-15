package main

import (
	"github.com/colmpat/words-that-matter/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	return NewAdminHandler(&Config{
		r.Group("/admin"),
	})
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	makeAuthHandler(r)
	makeAdminHandler(r)

	r.Run()
}
