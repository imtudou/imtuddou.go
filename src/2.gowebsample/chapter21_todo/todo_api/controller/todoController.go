package controller

import (
	"2.gowebsample/chapter21_todo/todo_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
url → controller → models → dao
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Create(c *gin.Context) {

	//1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	//2. 写入数据库
	err := models.Create(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "成功",
			"data": todo.ID,
		})
	}

}

func Update(c *gin.Context) {

	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	//1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	todo.ID = uint(idUint)

	//2. 写入数据库
	err := models.Update(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "成功",
			"data": todo.ID,
		})
	}
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, _ := strconv.ParseUint(id, 10, 64)
	if err := models.Delete(uint(idUint)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"data": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "成功",
			"data": true,
		})
	}
}

func GetList(c *gin.Context) {

	todos, err := models.GetList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
			"data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "成功",
			"data": todos,
		})
	}
}
