package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	defer fmt.Println("sum.n1 = ", n1)
	defer fmt.Println("sum.n2 = ", n2)
	res := n1 + n2
	fmt.Println("sum.res = ", res)
	return res
}

func main() {
	res := sum(11, 23)
	fmt.Println("main.res = ", res)

	// output;
	// sum.res =  34
	// sum.n2 =  23
	// sum.n1 =  11
	// main.res =  34
}
