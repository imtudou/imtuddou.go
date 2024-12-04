package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default() // 返回默认的路由引擎

	// 指定用户使用get 请求访问/hello时 执行 sayHello这个函数
	r.GET("/hello", sayHello)

	// 启动服务
	r.Run(":9090")
}

func sayHello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hello world",
	})
}
