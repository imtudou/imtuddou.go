package main

import (
	"fmt"
	student "sample/chapter9/model/student"
)

func main() {
	stu := student.Student{Name: "tom", Age: 19}
	fmt.Println(stu)

	stu2 := student.CreateStudent("tom", 19, 11.2, 190.9)
	fmt.Println(*stu2)

	fmt.Println(stu2.GetHeight())

}

// Golang的结构体没有构造函数,通常可以使用工厂模式来解决这个问题。
