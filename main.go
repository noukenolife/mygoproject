package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/noukenolife/mygoproject/helper"
	"github.com/noukenolife/mygoproject/infrastructure/di"
)

func main() {
	godotenv.Load()

	Port := helper.Getenv("PORT")
	AppName := helper.GetenvWithDefaultValue("APP_NAME", "mygoproject")
	AppSecret := helper.GetenvWithDefaultValue("APP_SECRET", "secret")

	r := gin.Default()

	store := sessions.NewCookieStore([]byte(AppSecret))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
	})

	r.Use(sessions.Sessions(AppName, store))

	c := di.CreateContainer()
	c.ApiRouter.InitRouter(r)
	c.WebRouter.InitRouter(r)

	if Port != "" {
		r.Run(fmt.Sprintf(":%s", Port))
	} else {
		r.Run("127.0.0.1:3000")
	}
}
