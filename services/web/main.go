package main

import (
	"github.com/colmpat/words-that-matter/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func registerAuthProviders(r *gin.Engine) {
	// google
	gp := auth.NewGoogleProvider()
	auth.RegisterProvider(r, gp, "google")

	// can add more providers here
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	registerAuthProviders(r)

	r.Run()
}
