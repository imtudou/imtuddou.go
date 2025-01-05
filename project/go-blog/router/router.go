package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {

	// 路由
	// 三种路由返回类型 1.页面(views) 2.数据（api） 3.静态资源

	/*1.HTML*/
	http.HandleFunc("/", views.HTML.IndexHandle)
	//处理登录 http://localhost:8080/login
	http.HandleFunc("/login", views.HTML.LogInHandle)
	//处理category 的路由 http://localhost:8080/c/1
	http.HandleFunc("/c/", views.HTML.CategoryHandle)
	//处理文章详情  http://localhost:8080/p/7.html
	http.HandleFunc("/p/", views.HTML.PostDetailHandle)
	//http://localhost:8080/writing
	http.HandleFunc("/writing", views.HTML.WritingHadle)
	// 归档 http://localhost:8080/pigeonhole
	http.HandleFunc("/pigeonhole", views.HTML.PigeonholeHandle)

	// 关于 http://localhost:8080/about
	http.HandleFunc("/about", views.HTML.AboutHandle)
	/*2.API*/
	//处理登录接口 http://localhost:8080/api/v1/login
	http.HandleFunc("/api/v1/login", api.API.LogIn)
	// 发布文章
	http.HandleFunc("/api/v1/post", api.API.Save)
	// 查看文章详情 http: //localhost:8080/api/v1/post/25
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	// 搜索框懒加载列表 /api/v1/post/search?val=g
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	// url 请求映射
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
