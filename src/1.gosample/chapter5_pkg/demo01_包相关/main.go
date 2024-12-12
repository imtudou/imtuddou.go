package main

import (
	"fmt"
	"sample/chapter5/utils"
	util "sample/chapter5/utils"
)

func main() {

	// var n1 float64 = 1.2
	// var n2 float64 = 3.2
	// var operator byte = '-'

	// var res = cal(n1, n2, operator)
	// var res = cal(1.2, 3.2, '+')
	fmt.Println(utils.Cal(1.2, 3.2, '+'))
	fmt.Println(utils.Now())
	fmt.Println(util.HeroName) // 给包取别名

	fmt.Println(getSum(1, 10))
	res1, res2 := getSumAndSub(10, 22)
	fmt.Printf("res1 = %v, res2 = %v", res1, res2)

	test(10)

}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func test(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		test(n)
	} else {
		fmt.Println(n)
	}
}
