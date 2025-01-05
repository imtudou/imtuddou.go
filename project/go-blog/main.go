package main

import (
	"go-blog/common"
	"go-blog/router"
	"go-blog/server"
)

func init() {
	// 初始化模板加载
	common.LoadTemplate()
}

func main() {
	//1. 程序入口 一个项目只能有一个入口
	//2. web程序

	// 路由
	router.Router()
	server.App.Run("", "8080")
}
