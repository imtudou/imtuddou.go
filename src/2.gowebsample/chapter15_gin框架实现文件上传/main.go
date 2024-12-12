package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("./upload", func(c *gin.Context) {

		//UploadFile(c)          // 单个文件
		UploadMultipartFile(c) // 多个文件
	})
	r.Run(":9090")

}

// 单个文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	// 上传文件到指定目录
	c.SaveUploadedFile(file, "./"+file.Filename)
	c.JSON(http.StatusOK, gin.H{
		"msg": file.Filename + "上传成功",
	})
}

// 多个文件
func UploadMultipartFile(c *gin.Context) {

	// 设置上传文件的目录
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	// 获取表单中的文件
	form, _ := c.MultipartForm()
	files := form.File["f1"]
	for _, file := range files {
		// 定义文件存储路径
		filePath := filepath.Join(uploadDir, file.Filename)

		// 将文件保存到指定路径
		c.SaveUploadedFile(file, filePath)
		fmt.Println("Uploaded file:", filePath)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%d个文件上传成功 :", len(files)),
	})
}
