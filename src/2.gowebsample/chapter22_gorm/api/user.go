package api

import (
	"2.gowebsample/chapter22_gorm/dao"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
)

var UserAPI = &User{}

type User struct {
}

func (*User) Save(c *gin.Context) {
	u := dao.User{
		UserName: "zzss" + strconv.Itoa(rand.Intn(10)),
		Passwd:   "zzss",
		Avatar:   "zzss",
	}

	data, err := dao.SaveUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
	return
}

func (*User) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data, err := dao.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
	return
}

func (*User) GetUserAll(c *gin.Context) {
	data, err := dao.GetUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
	return
}

func (*User) UpdateUser(c *gin.Context) {
	u := dao.User{
		Model: gorm.Model{
			ID: 12,
		},
		UserName: "lisi",
		Passwd:   "lisi",
		Avatar:   "lisi",
	}
	err := dao.UpdateUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
	return
}

func (*User) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	u := dao.User{
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	err := dao.DeleteUser(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
	return
}
