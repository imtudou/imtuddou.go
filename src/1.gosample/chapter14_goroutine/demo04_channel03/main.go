package main

import (
	"fmt"
	"time"
)

// channel 使用细节和注意事项
func main() {
	// test
	fmt.Println("main---------------------")
	ch := make(chan int, 20)
	resultChan := make(chan bool, 5)
	//1. channel申明只读
	go send(ch, resultChan)

	//2. channel申明只写
	go receive(ch, resultChan)

	// 主线程需要等一下 不然输出太快了
	time.Sleep(1000 * time.Second)
	var length = len(resultChan)
	if length == 2 {
		fmt.Printf("resultChan length is %d\n", length)
	}
	fmt.Println("main---------------------")
}

// ch chan <- int 只读操作
func send(ch chan<- int, resultChan chan bool) {
	fmt.Println("send---------------------")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	resultChan <- true
}

func receive(ch <-chan int, resultChan chan bool) {
	fmt.Println("receive---------------------")

	for v := range ch {
		fmt.Println("receive.ch = %v", v)
	}
	resultChan <- true
}
