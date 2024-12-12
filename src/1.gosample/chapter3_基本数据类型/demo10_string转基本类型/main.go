package main

import (
	"fmt"
	"strconv"
)

// go string转基本类型
func main() {

	// ParseBool  返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	var str1 string = "true"
	var res1 bool
	res1, _ = strconv.ParseBool(str1)
	fmt.Printf("type=%T res1=%v \n", res1, res1)

	var str2 string = "3.1425"
	var res2 float64
	res2, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("type=%T res2=%v \n", res2, res2)

	var str3 string = "123444"
	var res3 int64
	res3, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Printf("type=%T res3=%v \n", res3, res3)

	// 注意:如果转换失败返回默认值
	var str4 string = "hello"
	var res4 int64 = 11
	res4, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("type=%T res3=%v \n", res4, res4) // output:type=int64 res3=0

}
