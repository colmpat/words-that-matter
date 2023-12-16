package main

import "github.com/gin-gonic/gin"

// AdminHandler is a handler for admin pages
type AdminHandler struct {
	*Handler
	navLinks []Link
}

// NewAdminHandler creates a new AdminHandler
func NewAdminHandler(c *Config) *AdminHandler {
	navLinks := []Link{
		{"/", "Back Home"},
		{"/admin", "Admin"},
		{"/admin/publish", "Publish"},
	}

	h := AdminHandler{
		NewHandler(c),
		navLinks,
	}

	h.g.GET("/", func(c *gin.Context) {
		c.HTML(200, "admin.html", gin.H{
			"links":    h.navLinks,
			"signedIn": true,
		})
	})

	h.g.GET("/publish", func(c *gin.Context) {
		c.HTML(200, "publish.html", gin.H{
			"links":    h.navLinks,
			"signedIn": true,
		})
	})

	return &h
}
