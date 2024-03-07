package server

import (
	"dothattome/reqests"
	"github.com/gin-gonic/gin"
	"sync"
)

// StartServer is a function for starting server
func StartServer() {
	r := gin.Default()
	var wg sync.WaitGroup
	wg.Add(1)
	reqests.SetupPath(r)

	go func() {
		r.GET("/end", func(c *gin.Context) {
			wg.Done()
		})

		err := r.Run("127.0.0.1:8081") // listen and serve on 0.0.0.0:8080
		if err != nil {
			panic(err)
		}

	}()

	wg.Wait()

}
