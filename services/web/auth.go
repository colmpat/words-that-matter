package main

import (
	"github.com/colmpat/words-that-matter/internal/auth"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	*Handler
}

// creates a new AuthHandler and registers all routes
func NewAuthHandler(c *Config, providers ...auth.Provider) *AuthHandler {
	h := AuthHandler{
		NewHandler(c),
	}

	for _, p := range providers {
		g := h.g.Group(p.Name())
		auth.RegisterProvider(g, p)
	}

	return &h
}
