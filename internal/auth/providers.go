package auth

import (
	"net/http"

	"github.com/colmpat/words-that-matter/pkg/models"
	"github.com/gin-contrib/sessions"
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

// SetUser sets the user in the session.
func SetUser(c *gin.Context, user models.User) {
	session := sessions.Default(c)
	session.Set("userId", user.ID)
	session.Set("isAdmin", user.IsAdmin)
	session.Save()
}

// UserAuth is a middleware that checks if a user is logged in. If not, it redirects to the login page.
func UserAuth() gin.HandlerFunc {
	return UserAuthTo("/login")
}

// UserAuth is a middleware that checks if a user is logged in. If not, it redirects to the url provided.
func UserAuthTo(location string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("userId") == nil {
			c.Redirect(http.StatusFound, location)
			c.Abort()
			return
		}
		c.Next()
	}
}

// AdminAuth is a middleware that checks if a user is logged in and is an admin. It redirects to "/login" if not signed in and "/" if signed in but not an admin.
func AdminAuth() gin.HandlerFunc {
	return AdminAuthTo("/login", "/")
}

// AdminAuthTo is a middleware that checks if a user is logged in and is an admin. It redirects to the noUserLoc if not signed in and the notAdminLoc if signed in but not an admin.
func AdminAuthTo(noUserLoc, notAdminLoc string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("userId") == nil {
			c.Redirect(http.StatusFound, noUserLoc)
			c.Abort()
			return
		}
		if session.Get("isAdmin") != true {
			c.Redirect(http.StatusFound, notAdminLoc)
			c.Abort()
			return
		}
		c.Next()
	}
}
