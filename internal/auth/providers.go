package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Provider is an interface for an authentication provider. It provides a login handler and a callback handler.
type Provider interface {
	LoginHandler(c *gin.Context)
	CallbackHandler(c *gin.Context)
}

// RegisterProvider registers a provider's login and callback handlers on the given gin engine.
func RegisterProvider(r *gin.Engine, provider Provider, providerName string) {
	loginPath := fmt.Sprintf("auth/%s/login", providerName)
	callbackPath := fmt.Sprintf("auth/%s/callback", providerName)

	r.GET(loginPath, provider.LoginHandler)
	r.GET(callbackPath, provider.CallbackHandler)
}
