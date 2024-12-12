package main

import (
	"fmt"
)

/*
请完成goroutine和channel协同工作的案例，具体要求:
1)开启一个writeData协程,向管道intChan中写入50个整数.
2)开启一个readData协程，从管道intChan中读取writeData写入的数据。
3)注意: writeData和readDate操作的是同一个管道
4)主线程需要等待writeData和readDate协程都完成工作才能退出【管道】
*/
func main() {
	intChan := make(chan int, 50)
	resultChan := make(chan bool, 50)
	go writeData(intChan)
	go readData(intChan, resultChan) // 不取数据就会 error deadlock!

	for v := range resultChan {
		if v {
			fmt.Printf("执行结果 = %v \n", v)
		}
	}
}

func writeData(intChan chan int) {
	fmt.Printf("writeData 开始 --------------\n")
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Printf("writeData 结束 --------------\n")
}

func readData(intChan chan int, resultChan chan bool) {
	fmt.Printf("readData 开始 --------------\n")
	for v := range intChan {
		fmt.Printf("readData = %v \n", v)
	}

	resultChan <- true
	close(resultChan)
	fmt.Printf("readData 结束-------------- \n")
}
