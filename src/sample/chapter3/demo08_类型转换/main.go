package main

import (
	"fmt"
	"unsafe"
)

// go 基本数据类型转换
func main() {
	var n1 int32 = 100

	// i =>  float
	var n2 float32 = float32(n1)

	var n3 int8 = int8(n1)
	var n4 int64 = int64(n1) // 低精度 → 高精度
	fmt.Printf("n1=%v,n2=%v,n3=%v,n4=%v \n", n1, n2, n3, n4)
	fmt.Printf("n2=%v,n1数据类型:%T,n2数据类型:%T,n2占用字节:%d, \n", n2, n1, n2, unsafe.Sizeof(n2))

	// 转换的是变量存储的值 变量本身数据类型并没有改变
	fmt.Printf("n1 type is %T\n", n1) // int32

	// 在转换中比如将int64 转成 int8[-128 → 127] 编译时不会报错 只是转换时按照溢出处理 和我们希望的结果不一样
	var n5 int64 = 999999
	var n6 int8 = int8(n5)
	fmt.Println("n6=", n6)
}
