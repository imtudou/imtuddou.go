package main

import (
	"fmt"
	"strconv"
)

// go 基本数据类型转string
func main() {
	var n1 int32 = 100
	var n2 float32 = 12.345
	var n3 bool = true
	var n4 byte = 'h'
	 var str string 

	// 1. 使用fmt.Sprintf()转
	str = fmt.Sprintf("%d",n1)
	fmt.Printf("type=%T str=%q \n",str,str)

	str = fmt.Sprintf("%f",n2)
	fmt.Printf("type=%T str=%q \n",str,str)

	str = fmt.Sprintf("%t",n3)
	fmt.Printf("type=%T str=%q \n",str,str)

	str = fmt.Sprintf("%c",n4)
	fmt.Printf("type=%T str=%q \n",str,str)



	// 2. 使用 strconv.xxx()转
	fmt.Printf("----------------------------------\n")
	fmt.Printf("2. 使用 strconv.xxx()转------------\n")
	str = strconv.FormatInt(int64(n1),10)
	fmt.Printf("type=%T str=%q \n",str,str)

	str = strconv.FormatFloat(float64(n2),'f',10,64)
	fmt.Printf("type=%T str=%q \n",str,str)

	str = strconv.FormatBool(n3)
	fmt.Printf("type=%T str=%q \n",str,str)

	//Itoa是FormatInt(i, 10) 的简写。 直接将int 转string
	str = strconv.Itoa(int(n1))
	fmt.Printf("type=%T str=%q \n",str,str)
}
