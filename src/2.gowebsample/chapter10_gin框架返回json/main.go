package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		data := gin.H{
			"message": "pong",
			"name":    "张三",
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")

}
