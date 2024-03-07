package reqests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// CallTestGet is a test function if server working
func CallTestGet() {
	response, err := http.Get("http://127.0.0.1:8081/test")
	if err != nil {
		fmt.Println(err)
	}

	rdata, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(rdata))
}

// CallKillServer is a function for kiling server.
func CallKillServer() {
	_, err := http.Get("http://127.0.0.1:8081/end")
	if err != nil {
		fmt.Println(err)
	}

}
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
			"xxx":     "yyyy",
		})
	})

	ginEngine.GET("/bump", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "lol",
		})
	})
}
