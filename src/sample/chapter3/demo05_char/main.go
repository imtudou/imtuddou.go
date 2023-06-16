package main

import (
	"fmt"
)

// go 字符char→byte类型使用
func main() {

	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1, "c2=", c2) // 直接输出其实是输出了对应字符的 ascii 码 需要转义一下
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//var c3 byte = '中'
	//fmt.Printf("c3=%c", c3)// constant 20013 overflows byte

	var c3 int = '中'
	fmt.Printf("c3=%c  对应的码值=%d \n", c3, c3) // c3=中  对应的码值=20013

	// 20013 → 输出 '中'
	var c4 int = 20013
	fmt.Printf("c4=%c \n", c4) // c4=中

	// 字符类型是可以运算的 相当于一个整数 按照ascii码运行
	var c5 = 10 + 'a' // 相当于 10+97 = 107
	fmt.Println("c5=", c5)
}
