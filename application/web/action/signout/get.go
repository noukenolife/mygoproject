package signout

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type GetSignout struct{}

func (self GetSignout) Invoke(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/signin")
}
