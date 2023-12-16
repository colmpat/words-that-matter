package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/colmpat/words-that-matter/pkg/db"
	"github.com/colmpat/words-that-matter/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// GoogleProvider is an implementation of the Provider interface for Google OAuth.
type GoogleProvider struct {
	config           *oauth2.Config
	oauthStateString string
}

// NewGoogleProvider returns a new GoogleProvider. It reads the client ID, client secret, and redirect URL from environment variables.
func NewGoogleProvider() *GoogleProvider {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	// Random string for OAuth state
	oauthStateString := "random"

	return &GoogleProvider{
		config,
		oauthStateString,
	}
}

func (gp *GoogleProvider) Name() string {
	return "google"
}

func (gp *GoogleProvider) LoginHandler(c *gin.Context) {
	url := gp.config.AuthCodeURL(gp.oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (gp *GoogleProvider) CallbackHandler(c *gin.Context) {
	state := c.Query("state")
	if state != gp.oauthStateString {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid state"})
		return
	}

	code := c.Query("code")
	token, err := gp.config.Exchange(c, code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := gp.GetUserEmail(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	SetUser(c, user)
	c.Redirect(http.StatusFound, "/")
}

func (gp *GoogleProvider) GetUserEmail(token *oauth2.Token) (models.User, error) {
	client := gp.config.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return models.User{}, err
	}

	return db.ProdDB.GetUserByEmail(userInfo.Email)
}
