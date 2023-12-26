package main

import "fmt"

// 定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

// 上面的申明方式也可一次性申明
var (
	n3    = 300
	n4    = 400
	name2 = "mary"
)

func main() {

	// 变量声明

	var a = 1
	fmt.Println(a)

	var b = 2
	fmt.Println(b)

	// 演示一次性申明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 演示一次性申明多个变量
	// var n1, name, n3 = 100, "tom", 999
	// fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	/*
		等同于
		var name string
		name = "tom"
	*/
	// name := "tom"
	// fmt.Println("name=", name)

	fmt.Println("n1=", n1, "n2=", n2, "name=", name, "n3=", n3, "n4=", n4, "name2=", name2)

}
