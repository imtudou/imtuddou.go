package main

import "fmt"

// 一个被测试的函数
func add(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res

}

/*
测试的具体例子
*/
func main() {

	// 传统的测试方法就是在main函数里面调用一次看看结果是否符合预期
	res := add(10) // 6
	if res == 55 {
		fmt.Printf("符合预期：res = %v\n", res)
	} else {
		fmt.Printf("不符合预期：res = %v\n", res)
	}

}
