package main

import (
	"fmt"
	"strconv"
)

func add() func(int) int {
	var n = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += strconv.Itoa(x)
		fmt.Println(str)
		return n
	}
}

func main() {
	f := add()
	fmt.Println(f(1)) // output: 11
	fmt.Println(f(1)) // output: 12
	fmt.Println(f(1)) // output: 13
}
