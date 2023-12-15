package main

import "github.com/gin-gonic/gin"

// Config is the configuration for a Handler
type Config struct {
	g *gin.RouterGroup
}

// Handler is a handler for a portion of the API
type Handler struct {
	*Config
}

// NewHandler returns a new Handler
func NewHandler(config *Config) *Handler {
	return &Handler{config}
}
