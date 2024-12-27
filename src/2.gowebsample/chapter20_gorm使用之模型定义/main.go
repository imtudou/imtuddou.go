package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

// 定义模型
type UserTest struct {
	gorm.Model // 内嵌gorm.Model
	Name       string
	Age        sql.NullInt64 `gorm:"column:user_age"` // 零值类型
	Birthday   *time.Time
	Email      string  `gorm:"type:varchar(100);unique_index"`
	Role       string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNo   *string `gorm:"unique;not null"` // 设置会员卡号唯一且不为空
	Num        int     `gorm:"AUTO_INCREMENT"`  // 设置num自增
	Address    string  `gorm:"index:addr"`      // 给 address 字段创建名为 addr 的索引
	IgnoreMe   int     `gorm:"-"`               // 忽略字段
}

// 使用 AnimalID 作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"` // 主键
	Name     string
	Age      sql.NullInt64 `gorm:"column:animal_age"`
}

// 指定表名
func main() {

	// 连接mysql数据库
	dsn := "root:12345@tcp(127.0.0.1:3306)/finbook_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // 按照默认命名规则，数据库中的表名应该是 users。但是，启用这个前缀后，表名会变成 t_users。
			SingularTable: true,                              // 禁用复数 如果启用这个选项，并且你的结构体名为 User，即使按照常规规则应该是复数形式，表名也会保持为单数形式 user（但注意，由于同时设置了 TablePrefix，最终表名将是 t_user）。
			NoLowerCase:   true,                              // 默认情况下，GORM 会将 Go 语言中的驼峰命名（如 UserName）转换为小写加下划线（user_name）的格式。如果启用这个选项，名称将保持原样（但注意，由于同时设置了 TablePrefix 和 SingularTable，这些选项仍然会影响最终表名）。
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // 这是一个名称替换器，用于在将结构体或字段名转换为数据库名之前进行字符串替换。在这个例子中，所有的 CID 都会被替换为 Cid。这意味着，如果你的结构体中有一个字段名为 CID，在数据库中这个字段将会以 Cid 的形式出现。
		},
		NowFunc: func() time.Time {
			return time.Now().Local() // 更改创建时间使用的函数
		},
	})
	if err != nil {
		panic(err)
	}

	// 迁移
	err = db.AutoMigrate(&UserTest{})
	if err != nil {
		panic(err)
		return
	}
	err = db.AutoMigrate(&Animal{})
	if err != nil {
		return
	}

	//使用 usertest 结构体创建 名为  UserTest_P1 的表 适合创建分表时使用
	db.Table("UserTest_P1").AutoMigrate(&UserTest{})

	// 创建记录
	a := Animal{
		AnimalID: 0,
		Name:     "张阿三",
		Age:      sql.NullInt64{20, true},
	}
	result := db.Create(&a) // 通过数据的指针 写入数据
	fmt.Printf("AnimalID = %v, Error = %v, RowsAffected = %v  \r\n", a.AnimalID, result.Error, result.RowsAffected)

	// 创建多条记录
	a1 := []*Animal{
		&Animal{
			AnimalID: 0,
			Name:     "张阿四",
			Age:      sql.NullInt64{20, true},
		},
		&Animal{
			AnimalID: 0,
			Name:     "张阿五",
			Age:      sql.NullInt64{20, true},
		},
	}
	result = db.Create(&a1) // 通过数据的指针 写入数据
	fmt.Printf("AnimalID = %v, Error = %v, RowsAffected = %v  \r\n", a.AnimalID, result.Error, result.RowsAffected)

	// 指定字段创建
	db.Select("Name", "Age").Create(&a)

}
