package main

import (
	"github.com/gin-gonic/gin"
	"server/reqests"
)

func main() {
	r := gin.Default()

	reqests.SetupPath(r)

	err := r.Run("127.0.0.1:8081") // listen and serve on 0.0.0.0:8080

	if err != nil {
		panic(err)
	}
}
