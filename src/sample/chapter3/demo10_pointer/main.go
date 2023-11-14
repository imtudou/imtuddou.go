package main

import (
	"fmt"
	"D:/software/go/src/src/sample/chapter3/model/model"
)

// go 指针的使用
func main() {

	var n1 int = 10
	fmt.Println("\n n1 = ", n1)
	fmt.Println("n1的地址:", &n1)

	var ptr *int
	ptr = &n1
	*ptr = 11 // 这里赋值时 n1 的值会改变(指针改变)
	fmt.Println("\n n1 = ", n1)
	fmt.Println("ptr的地址:", ptr)

	// 可以 但不推荐
	var int int = 111
	fmt.Println("\n int = ", int)

	fmt.Println(model.HeroName)
}
