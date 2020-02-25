package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/http"
)

func main() {
	r := gin.Default()
	http.InitRouter(r)
	r.Run("127.0.0.1:8080")
}
