package main

import (
	"fmt"
	"unsafe"
)

// go 基本数据类型转换
func main() {
	var n1 int32 = 100
	var n2 float32 = float32(n1)
	fmt.Printf("n1=%v,n2=%v \n", n1, n2)
	fmt.Printf("n2=%v,n1数据类型:%T,n2数据类型:%T,n2占用字节:%d,", n2, n1, n2, unsafe.Sizeof(n2))

	// 转换的是变量存储的值 变量本身数据类型并没有改变

}
