package eventsub

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscribeEvent struct {
	Token     string `json:"token" binding:"required"`
	Challenge string `json:"challenge" binding:"required"`
	Type      string `json:"type" binding:"required"`
}

func Post(c *gin.Context) {
	var request SubscribeEvent
	if err := c.ShouldBindJSON(&request); err != nil {
		c.String(http.StatusBadRequest, "Invalid event subscription request.")
		return
	}

	c.String(http.StatusOK, request.Challenge)
}
