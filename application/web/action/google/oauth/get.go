package oauth

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/noukenolife/mygoproject/application/service/authenticate_service"
)

type GetGoogleOAuth struct {
	GetAuthURL authenticate_service.GetAuthUrl
}

func (self GetGoogleOAuth) Invoke(c *gin.Context) {
	session := sessions.Default(c)

	state := uuid.New().String()

	session.Clear()
	session.Set("state", state)
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	authURL := self.GetAuthURL.Invoke(state)
	c.Redirect(http.StatusTemporaryRedirect, authURL.Value)
}
