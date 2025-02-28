package main

import (
	"github.com/gin-gonic/gin"
)

type TestParams struct {
	Name       string            `form:"name" json:"name" binding:"required"`
	Age        int               `form:"age" json:"age" binding:"required"`
	Addr       []string          `form:"address" json:"addr" binding:"required"`
	AddressMap map[string]string `form:"address" json:"addressMap"`
}

// 获取 query 提交的参数
// GET请求 URL ?后面的是querystring参数
// key=value格式，多个key-value用 & 连接

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {

		// eq:  http://localhost:9090/index?name=zs&age=23&address=123
		// 1.获取querystring参数
		//name := c.Query("name")
		//age := c.Query("age")
		//address, ok := c.GetQuery("address")         // 获取一个参数是否存在
		//address1 := c.DefaultQuery("address", "默认值") // 如果这个参数不存在 就给个默认值
		//c.JSON(http.StatusOK, gin.H{
		//	"name":     name,
		//	"age":      age,
		//	"address":  address,
		//	"ok":       ok,
		//	"address1": address1,
		//})

		//2. 定义结构体直接获取querystring 参数
		//var test TestParams
		////c.BindQuery(&test) // 如果添加了 require tag 使用这个的时候 返回的状态码会变成400：参数错误 结果正常返回
		//c.ShouldBindQuery(&test)
		//c.JSON(http.StatusOK, test)

		//3.接收数组参数
		//// eq:  http://localhost:9090/index?name=zs&age=23&address=123&address=456&address=789
		//address := c.QueryArray("address")
		//c.JSON(200, gin.H{"address": address})

		//3.1 接收数组参数 绑定结构体
		//var test TestParams
		//c.ShouldBindQuery(&test)
		//c.JSON(200, test)

		//4.接受结构体
		//// eq:  http://localhost:9090/index?name=zs&age=23&address=123&address=456&address=789
		var test TestParams
		c.ShouldBindQuery(&test)
		test.AddressMap = c.QueryMap("addressMap") // 需要单独赋值
		c.JSON(200, test)

	})
	r.Run(":9090")
}
