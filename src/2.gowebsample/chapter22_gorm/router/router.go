package router

import (
	"2.gowebsample/chapter22_gorm/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/save", api.UserAPI.Save)
	r.GET("/getuser", api.UserAPI.GetUser)
	r.GET("/getuserall", api.UserAPI.GetUserAll)
	r.GET("/update", api.UserAPI.UpdateUser)
	r.GET("/delete", api.UserAPI.DeleteUser)
}
