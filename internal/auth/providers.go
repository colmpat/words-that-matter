package auth

import (
	"github.com/gin-gonic/gin"
)

// Provider is an interface for an authentication provider. It provides a login handler and a callback handler.
type Provider interface {
	LoginHandler(c *gin.Context)
	CallbackHandler(c *gin.Context)
	Name() string
}

// RegisterProvider registers a provider's login and callback handlers on the given gin engine.
func RegisterProvider(g *gin.RouterGroup, provider Provider) {
	g.GET("login", provider.LoginHandler)
	g.GET("callback", provider.CallbackHandler)
}
