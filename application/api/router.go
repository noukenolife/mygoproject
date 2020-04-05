package api

import (
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/api/action/helloworld"
	"github.com/noukenolife/mygoproject/application/middleware/checkauth"
)

type ApiRouter struct {
	GetHelloWorld  helloworld.GetHelloWorld
	PostHelloWorld helloworld.PostHelloWorld
}

func (self *ApiRouter) InitRouter(r *gin.Engine) {
	api := r.Group("/api", checkauth.CheckAuth())
	api.GET("/hello", self.GetHelloWorld.Invoke)
	api.POST("/hello", self.PostHelloWorld.Invoke)
}
