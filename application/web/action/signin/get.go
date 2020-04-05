package signin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetSignin struct{}

func (self GetSignin) Invoke(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{})
}
