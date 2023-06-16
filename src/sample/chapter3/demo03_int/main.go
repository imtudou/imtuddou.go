package main

import (
	"fmt"
	"unsafe"
)

// 整数int的基本使用

/*
整数类型
类型	存储	范围
int8	1字节 	-128~127
int16	2字节	-2^15~2^15-1
int32	4字节	-2^31~2^31-1
int64	8字节	-2^63~2^63-1

整型的类型

类型	有无符号	占用存储空间				表数范围	备注
int		有			32位系统4个字节64位系统8个字节	-231 ~ 231-1-263 ~ 263-1uint无32位系统4个字节64位系统8个字节0~232-10~264-1rune有与int32一样-231 ~ 231-1等价int32,表示一个 Unicode 码byte无与 uint8 等价0~255当要存储字符时选用byte

*/

func main() {
	var i int8 = 127
	//var i1 int8 = 128 // constant 128 overflows int8
	fmt.Println("i=", i)
	//fmt.Println("i1=", i1)

	// uint8(0~255) 其他的 uint16 uint32 uint64类型推断即可
	var u uint16 = 255
	fmt.Println("u=", u)

	// int uint rune byte的使用
	var a int = 11111111
	var b uint = 1
	var c byte = 223

	fmt.Println("a=", a, "b=", b, "c=", c)

	// 整型使用细节
	var n1 = 100 // n1 是什么类型 int  uint16?  使用fmt.Printf()格式化输出
	fmt.Printf("n1的类型 %T \n", n1)

	// 如何在程序中查看某个变量占用的字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2的类型 %T   n2占用的字节 %d", n2, unsafe.Sizeof(n2))

	// Golang程序中整型变量在使用时，遵守保小不保大的原则，即:在保证程序正,确运行下，尽量使用占用空间小的数据类型。【如：年龄】
	var age byte = 120 // byte 0~255
	// bit：计算机中的最小存储单位。byte：计算机中基本存储单元。1byte = 8 bit
}
