package main

import "fmt"

// 1. 在Go中,函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量了。通过该变量可以对函数调用。
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 2. 函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用！
func myFun(funvar func(int, int) int, n1 int, n2 int) int {
	return funvar(n1, n2)
}

// 3. 为了简化数据类型定义, Go支持自定义数据类型基本语法: type自定义数据类型名数据类型//理解:相当于一个别名案例：type mylnt int// 这时 mylnt 就等价 int 来使用了.案例： type mySum func (int, int) int // 这时 mySum 就等价 一个 函数类型 func (int, int) int
type myFuntType func(int, int) int

func myFun3(funvar myFuntType, n1 int, n2 int) int {
	return funvar(n1, n2)
}

// 4. 支持對匿名函數返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 5. 可变参数: 编写一个函数sum ，可以求出 1到多个int的和(可变参数需要放到形参后面)
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// 函数调用，函数注意事项
func main() {
	res1 := getSum(10, 20)
	fmt.Print("\n res1=", res1)

	res2 := myFun(getSum, 20, 30)
	fmt.Print("\n res2=", res2)

	res3 := myFun3(getSum, 120, 130)
	fmt.Print("\n res3=", res3)

	res4, res5 := getSumAndSub(100, 29)
	fmt.Printf("\n res4=%v,res5=%v,", res4, res5)

	// 使用_标识符，忽略返回值
	// res4, _ := getSumAndSub(100, 29)
	// fmt.Printf("\n res4=%v", res4)

	res6 := sum(1, 10, 20, 30, 40, 50, 1)
	fmt.Print("\n res6=", res6)

}
