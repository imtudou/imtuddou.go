package main

import (
	"fmt"
)

func main() {

	num := 100
	fmt.Printf("num 类型=%T,值=%v,地址=%v \n", num, num, num)

	// new() 给值类型分配内存, make 给引用类型分配内存 做了两件事
	// 1. 分配空间装了个数字
	// 2. 在分配一个空间把地址放进去在让num2指向这个空间
	num2 := new(int)
	*num2 = 100
	fmt.Printf("num2 类型=%T,值=%v,地址=%v \n", num2, num2, num2)
}
