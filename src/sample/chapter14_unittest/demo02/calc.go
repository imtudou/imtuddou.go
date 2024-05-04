package demo02

// 一个被测试的函数
func add(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res

}

func sub(n1 int, n2 int) int {
	return n1 - n2
}
