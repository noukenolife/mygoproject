package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/noukenolife/mygoproject/application/http"
)

func main() {
	PORT := os.Getenv("PORT")
	r := gin.Default()
	http.InitRouter(r)

	if PORT != "" {
		r.Run(fmt.Sprintf(":%s", PORT))
	} else {
		r.Run("127.0.0.1:8080")
	}
}
