package main

import (
	"fmt"
)

func main() {
	fibonacci(1, 1)
}

// 斐波那契
func fibonacci(s int, e int) {
	fmt.Println(s)
	if e <= 1000 {
		sum := s + e
		fibonacci(e, sum)
	}
}
