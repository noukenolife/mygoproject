package checkauth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/dto/user_dto"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userProfileJSON := session.Get("USER")
		if userProfileJSON == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/signin")
			c.Abort()
			return
		}

		var userProfile user_dto.UserProfile
		json.Unmarshal([]byte(userProfileJSON.(string)), &userProfile)
		c.Set("USER", userProfile)

		c.Next()
	}
}
