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
	fmt.Printf("n2=%v,n1数据类型:%T,n2数据类型:%T,n2占用字节:%d,\n", n2, n1, n2, unsafe.Sizeof(n2))

	// 转换的是变量存储的值 变量本身数据类型并没有改变

	//在转换中，比如将int64 转成int8 【-128---127】 ，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样
	var n3 int64 = 9999999
	var n4 int8 = int8(n3);
	fmt.Println("n4=",n4) // n4= 127

}
