package main

import "fmt"

// 定义一个自定义的枚举类型
type Color int

// 定义枚举类型的常量
const (
	Red Color = iota
	Green
	Blue
)

func main() {
	// 直接使用Color.Red
	var color Color = Red
	fmt.Println("The color is:", color) // 输出: The color is: 0

	// 在switch语句中使用Color.Red
	switch color {
	case Red:

		fmt.Println("The color is red.")
	case Green:

		fmt.Println("The color is green.")
	case Blue:
		fmt.Println("The color is blue.")
	default:
		fmt.Println("Unknown color.")
	}

}
