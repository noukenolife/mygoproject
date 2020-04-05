package helloworld

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetHelloWorld struct{}

func (self GetHelloWorld) Invoke(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
