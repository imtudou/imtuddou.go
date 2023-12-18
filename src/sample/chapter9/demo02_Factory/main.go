package main

import (
	"fmt"
)

func main() {
	stu := model.Student{Name: "tom", Age: 19}
	fmt.Println(stu)

}

// Golang的结构体没有构造函数,通常可以使用工厂模式来解决这个问题。
