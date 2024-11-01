package main

import (
	"fmt"
	"time"
)

// goroutine 中使用recover 解决协程中出现panic 导致程序崩溃问题
/*
说明:如果我们起了一个协程,但是这个协程出现了panic,如果我们没有捕获这个panic,
就会造成整个程序崩溃，这时我们可以在goroutine中使用recover来捕获panic,
进行处理,这样即使这个协程发生的问题,但是主var mapl map[stringlstring线程仍然不受影响，可以继续执行
*/
func main() {
	go sayhello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ---------- ", i)
		time.Sleep(time.Second)

	}
}

func sayhello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("hello world %v \n", i)
	}
}

func test() {
	//	使用defer+recover
	defer func() {
		//捕获 test 抛出的 panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	//	定义一个map
	var errMap map[int]string
	errMap[0] = "捕获到异常信息"
}
