package main

import (
	"2.gowebsample/chapter21_todo/todo_api/dao"
	"2.gowebsample/chapter21_todo/todo_api/models"
	"2.gowebsample/chapter21_todo/todo_api/routers"
)

func main() {

	err := dao.InitMySQL()
	//1. 创建数据库 database
	//2. 连接数据库
	if err != nil {
		panic(err)
	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetRoute()
	r.Run(":9090")
}
