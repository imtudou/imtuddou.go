package routers

import (
	"2.gowebsample/chapter21_todo/todo_api/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute() *gin.Engine {

	r := gin.Default()
	// 静态文件
	//r.Static("/static", "./static")
	// 模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1group := r.Group("api/v1")
	{
		// 待办
		v1group.GET("/todo", controller.GetList)
		v1group.POST("/todo", controller.Create)
		v1group.PUT("/todo/:id", controller.Update)
		v1group.DELETE("/todo/:id", controller.Delete)
	}
	return r
}
