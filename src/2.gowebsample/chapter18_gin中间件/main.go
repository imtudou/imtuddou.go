package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/*
Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
*/
func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 1.注册一个全局中间件
	r.Use(TimeMiddleware(), m1())

	// 2.单独注册一个中间件给这个路由使用
	r.GET("/test", m2)

	// 3. 为路由组注册中间件
	// 写法1：
	//shopGroup := r.Group("/shop", m3)
	//{
	//	shopGroup.GET("/item", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"msg": c.Request.URL.Path,
	//		})
	//	})
	//}

	// 写法2：
	shopGroup := r.Group("/shop")
	shopGroup.Use(m3)
	{
		shopGroup.GET("/item", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": c.Request.URL.Path,
			})
		})

	}

	r.Run(":9090")

}

// 记录接口耗时的中间件：全局使用
func TimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()

		//此处可写相关判断逻辑
		c.Next() // 调用后续的处理函数

		//c.Abort() //阻止调用后续的处理函数
		ts := time.Since(begin)
		fmt.Printf("接口：%v, 耗时：%v  \r\n", c.Request.URL.Path, ts)
	}
}

// 中间件赋值：全局使用
func m1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("name", "李茶茶") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	}
}

// 为某个路由单独注册
func m2(c *gin.Context) {
	name, exists := c.Get("name") // 冲上下文取值
	if !exists {
		name = "匿名用户"
	}

	c.JSON(200, gin.H{
		"name": name,
	})
}

// 为路由组注册中间件
func m3(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "为路由组注册中间件",
	})
}
