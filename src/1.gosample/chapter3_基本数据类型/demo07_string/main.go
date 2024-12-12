package main

import (
	"fmt"
)

// go string 类型使用
func main() {

	c1 := "hello world"
	fmt.Println(c1)
	c1 = "111"
	fmt.Println(c1)

	// 字符串数组不可赋值
	//c2[0] = "a"

	// 字符串输出代码块
	c2 := `
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, World!")
		fmt.Println("111")
		fmt.Println("hahahah\r\nhahah")
		fmt.Println("c:\\d\\cdd\\ccc.txt")
	}
	`
	fmt.Println(c2)

	// 换行
	c3 := "hello\nworld"
	fmt.Println(c3)

	// 字符串拼接
	c4 := "hello" + "world"
	c4 += "hhahha"
	fmt.Println(c4)

	// 字符串换行拼接 + 必须放后面  因为go 默认会在结束语句后加 ;
	c5 := "hello" + "world" +
		"hello" + "world" +
		"hello" + "world" +
		"hello" + "world" +
		"hello" + "world" +
		"hello" + "world" +
		"hello" + "world"
	fmt.Println(c5)

	// 默认值
	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Println(a, b, c, d, e)                            // 定义类型不赋值有默认值
	fmt.Printf("a=%d,b=%v,c=%v,d=%v,e=%v", a, b, c, d, e) // 定义类型不赋值有默认值 a=0,b=0,c=0,d=false,e=

}
