package main

import (
	"fmt"
	"sample/chapter5/utils"
)

func main() {

	// var n1 float64 = 1.2
	// var n2 float64 = 3.2
	// var operator byte = '-'

	// var res = cal(n1, n2, operator)
	// var res = cal(1.2, 3.2, '+')
	fmt.Println(utils.Cal(1.2, 3.2, '+'))
	fmt.Println(utils.Now())
}
