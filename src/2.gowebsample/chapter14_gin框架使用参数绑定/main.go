package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
form 标签用于指定表单字段的名称
json 标签用于 JSON 数据的绑定
binding:"required" 表示这些字段是必填的
*/
type UserInfo struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

/*ShouldBind会按照下面的顺序解析请求中的数据完成绑定：

如果是 GET 请求，只使用 Form 绑定引擎（query）。
如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）
*/

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 绑定QueryString示例 (/LoginQuery?name=admin&pwd=123456)
	r.GET("/LoginQuery", func(c *gin.Context) {
		// 申明一个userinfo 类型的变量
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": user.Name,
				"pwd":  user.Pwd,
			})
		}

	})

	// 绑定form表单示例 (name=admin&pwd=123456)
	r.POST("/LoginForm", func(c *gin.Context) {
		// 申明一个userinfo 类型的变量
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": user.Name,
				"pwd":  user.Pwd,
			})
		}

	})

	// 绑定JSON的示例 ({"name": "admin", "pwd": "123456"})
	r.POST("/LoginJson", func(c *gin.Context) {
		// 申明一个userinfo 类型的变量
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name": user.Name,
				"pwd":  user.Pwd,
			})
		}

	})

	r.Run(":9090")
}
