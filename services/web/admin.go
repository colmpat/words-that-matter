package main

import "github.com/gin-gonic/gin"

// AdminHandler is a handler for admin pages
type AdminHandler struct {
	*Handler
}

// NewAdminHandler creates a new AdminHandler
func NewAdminHandler(c *Config) *AdminHandler {
	h := AdminHandler{
		NewHandler(c),
	}

	h.g.GET("/", func(c *gin.Context) {
		c.HTML(200, "admin.html", gin.H{})
	})

	h.g.GET("/publish", func(c *gin.Context) {
		c.HTML(200, "publish.html", gin.H{})
	})

	return &h
}
