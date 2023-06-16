package main

import (
	"fmt"
)

// go 小数float类型使用
func main() {
	var num1 float32 = 1.65
	var num2 float32 = -1.000600005
	var num3 float64 = -12564256.000600005
	fmt.Println("num1=", num1, "num2=", num2, "num3=", num3)

	// 精度丢失
	var num4 float32 = -1.000600005
	var num5 float64 = -1.000600005
	fmt.Println("num4=", num4, "num5=", num5) // num4= -1.0006 num5= -1.000600005

	// 浮点数默认类型 float64
	var num6 = 1.000600005
	fmt.Printf("num6类型 %T  \n", num6) //num6类型 float64

	// 十进制形式：5.12    .512
	num7 := 5.12
	num8 := .123 // output:0.123
	fmt.Println("num7=", num7, "num8=", num8)

	// 科学计数法
	num9 := 5.123e2                                              //5.123*10 的2次方
	num10 := 5.123e2                                             //5.123*10 的2次方
	num11 := 5.123e-2                                            //5.123 / 10 的2次方
	fmt.Println("num9=", num9, "num10=", num10, "num11=", num11) // output:num9= 512.3 num10= 512.3 num11= 0.05123
}
