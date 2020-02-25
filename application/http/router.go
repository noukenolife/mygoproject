package http

import (
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/http/action/helloworld"
	"github.com/noukenolife/mygoproject/application/http/action/slack/eventsub"
)

func InitRouter(r *gin.Engine) {
	r.GET("/hello", helloworld.Get)
	r.POST("/hello", helloworld.Post)

	r.POST("/slack/eventsub", eventsub.Post)
}
