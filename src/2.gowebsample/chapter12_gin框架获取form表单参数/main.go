package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"Msg": "1",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数
		name := c.PostForm("username")
		pwd := c.PostForm("password")

		// 取不到就返回默认值
		//name = c.DefaultPostForm("username","李茶茶")
		//pwd = c.DefaultPostForm("password","123")

		// 假装校验一下
		if name == "admin" && pwd == "123456" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Name": name,
				"Pwd":  pwd,
			})
		} else {
			//  账号或密码错误
			c.HTML(http.StatusOK, "login.html", gin.H{
				"Msg": "账号或密码错误！",
			})
		}

	})
	r.Run(":9090")

}
