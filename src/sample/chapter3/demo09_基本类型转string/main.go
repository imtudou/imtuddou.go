package main

import (
	"fmt"
	"strconv"
)

// go 基本数据类型转string
func main() {

	var n1 int = 99
	var n2 float64 = 23.4836
	var b bool = true
	var char byte = 'h'
	var str string

	// 使用 fmt.Sprintf 方法转换

	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type: %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("str type: %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type: %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type: %T str=%q\n", str, str)

	// 使用 strconv 方法转换
	var n3 int = 99
	var n4 float64 = 25.6985
	var b2 bool = true

	str = strconv.FormatInt(int64(n3), 10)
	fmt.Printf("str type: %T str=%q\n", str, str)

	str = strconv.FormatFloat(n4, 'f', 10, 64) // https://studygolang.com/pkgdoc
	fmt.Printf("str type: %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type: %T str=%q\n", str, str)

	// 注意: 如果类型转换失败为默认值而不是报错
	var str1 string = "hello"
	var n5 int64 = 11
	n5, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("n5 type: %T n5=%v\n", n5, n5)
}
