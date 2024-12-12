package main

import (
	"fmt"
)

// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
func sum(n1 int, n2 int) int {
	defer fmt.Println("sum.n1 = ", n1)
	defer fmt.Println("sum.n2 = ", n2)
	res := n1 + n2
	fmt.Println("sum.res = ", res)
	return res
}

func main() {
	res := sum(11, 23)
	fmt.Println("main.res = ", res)

	// output;
	// sum.res =  34
	// sum.n2 =  23
	// sum.n1 =  11
	// main.res =  34
}

// defer 注意事项和细节
// 1)值类型:基本数据类型int系列, float系列,bool, string 、数组和结构体struct
// 2)引用类型：指针、slice 切片、map、管道 chan、 interface 等都是引用类型
