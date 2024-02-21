package reqests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupPath(ginEngine *gin.Engine) {

	ginEngine.GET("/ping", func(c *gin.Context) {
		entries, err := os.ReadDir("./..")
		if err != nil {
			c.JSON(200, gin.H{
				"message": "error",
			})
		}
		for _, entry := range entries {
			fmt.Println(entry.Name())
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginEngine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "start",
		})
	})

	ginEngine.GET("/bump", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "lol",
		})
	})
}
