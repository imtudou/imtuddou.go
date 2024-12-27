package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*
https://gorm.io/zh_CN/docs/index.html
*/

type Product struct {
	gorm.Model // 内嵌gorm.Model
	Code       string
	Price      uint
	Stock      uint
}

func main() {
	// 连接mysql数据库
	dsn := "root:12345@tcp(127.0.0.1:3306)/finbook_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "789", Price: 200, Stock: 10000})

	// Read:拼接查询条件
	var product Product
	//db.First(&product, 2)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "456") // // 查找 code 字段值为 456 的记录
	//
	//sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	//	return tx.First(&product).Where("id = ?", 100).Limit(10).Order("age desc")
	//})
	// 生成sql：SELECT * FROM `products` WHERE code = '456' AND `products`.`deleted_at` IS NULL AND `products`.`id` = 1 ORDER BY `products`.`id` LIMIT 1
	//fmt.Printf("生成的sql：%v \r\n", sql)

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Where("id = ?", 1).Update("Price", 1000)
	//db.Model(&product).Where("id in ?", []int{1, 2, 3}).Update("Price", 1000)
	// Update - 更新多个字段
	//db.Model(&product).Where("id = ?", 1).Updates(Product{Price: 600, Stock: 1200})
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "456"})

	db.Delete(&product, 3) // 根据主键删除
	//fmt.Println(product)
	fmt.Println("执行完成！")
}
