package main

import (
	"github.com/colmpat/words-that-matter/pkg/db"
	"github.com/gin-gonic/gin"
)

// AdminHandler is a handler for admin pages
type AdminHandler struct {
	*Handler
	navLinks []Link
}

// NewAdminHandler creates a new AdminHandler
func NewAdminHandler(c *Config) *AdminHandler {
	navLinks := []Link{
		{"/admin", "Admin"},
		{"/admin/publish", "Publish"},
	}

	h := AdminHandler{
		NewHandler(c),
		navLinks,
	}

	h.g.GET("/", func(c *gin.Context) {
		c.HTML(200, "admin.html", gin.H{
			"links": navLinks,
		})
	})

	h.g.GET("/publish", func(c *gin.Context) {
		media, err := db.StagingDB.GetMedia()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.HTML(200, "publish.html", gin.H{
			"links": navLinks,
			"Media": media,
		})
	})

	return &h
}
