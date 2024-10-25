package main

import (
	"fmt"
	"runtime"
)

/*
1) go1.8后,默认让程序运行在多个核上,可以不用设置了
2) go1.8前,还是要设置一下,可以更高效的利益cpu
*/
func main() {
	num := runtime.NumCPU() // 获取当前系统的cpu数量

	runtime.GOMAXPROCS(num)
	fmt.Println("num = ", num)

}
