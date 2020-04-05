package oauth_callback

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/service/authenticate_service"
	"github.com/noukenolife/mygoproject/application/service/user_service"
)

type GetGoogleOAuthCallback struct {
	GetAccessToken      authenticate_service.GetAccessToken
	RetrieveUserProfile user_service.RetrieveUserProfile
}

func (self GetGoogleOAuthCallback) Invoke(c *gin.Context) {
	session := sessions.Default(c)

	state := session.Get("state")
	if state != c.Query("state") {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	code := c.Query("code")
	if code == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, err := self.GetAccessToken.Invoke(code)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userProfile, err := self.RetrieveUserProfile.Invoke(accessToken)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userProfileJSON, err := json.Marshal(userProfile)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Set("USER", string(userProfileJSON))
	err = session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
