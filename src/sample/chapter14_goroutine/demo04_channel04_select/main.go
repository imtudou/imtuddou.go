package main

import (
	"fmt"
	"time"
)

// 使用select 可以解决从管道中取数据的阻塞问题
func main() {

	// 1. 定义一个管道 10个 int类型
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2.定义一个管道 5个 string 列席
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%v", i)
	}

	// 传统方法遍历管道 如果不关闭会阻塞 导致  deadlock
	// 在实际开发中我们不确定什么时候应该关闭管道 可以使用select 方式解决
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intChan = %v \n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("stringChan = %v \n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("未匹配 \n")
			time.Sleep(time.Second)
			return
		}

	}

}
