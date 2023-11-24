package main

import "fmt"

var (
	// Func3 是一个全局匿名函数
	Func3 = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {

	//匿名函数1
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("匿名函数1:res1= ", res1)

	//匿名函数2
	func2 := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := func2(12, 22)
	fmt.Println("匿名函数2:res2=", res2)

	//全局匿名函数
	res3 := Func3(12, 22)
	fmt.Println("全局匿名函数:res3=", res3)

}
