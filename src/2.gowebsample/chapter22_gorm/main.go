package main

import (
	"2.gowebsample/chapter22_gorm/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
