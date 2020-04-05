package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/dto/user_dto"
)

type GetIndex struct{}

func (self GetIndex) Invoke(c *gin.Context) {
	userProfile := c.Keys["USER"].(user_dto.UserProfile)

	c.JSON(http.StatusOK, gin.H{
		"result": "authenticated",
		"email":  userProfile.Email.Value,
	})
}
