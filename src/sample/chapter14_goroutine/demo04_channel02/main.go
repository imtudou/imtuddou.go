package main

import "fmt"

/*
统计1~8000的数字中哪些是素数
*/
func main() {
	fmt.Printf("----------main begin--------\n")
	intChan := make(chan int, 10000)
	tempChan := make(chan int, 10000)
	resultChan := make(chan bool, 10)

	// 开启一个协程 向intChan 放入
	go writeData(intChan)

	// 开启4个协程读取数据
	for i := 0; i < 4; i++ {
		go readData(intChan, tempChan, resultChan)
	}

	length := len(resultChan)
	if length == 4 {
		// 当从resultChan 取出了4个结果就可以关闭 tempChan了
		close(tempChan)
	}
	fmt.Printf("length = %v \n", length)

	// 取十个出来看一下
	/*	for i := 0; i < 10; i++ {
		num := <-tempChan
		fmt.Printf("取%v个出来看一下= %v \n", i, num)
	}*/

	for v := range tempChan {
		fmt.Printf("取出来看一下= %v \n", v)
	}

	fmt.Printf("----------main end--------\n")

}

func writeData(intChan chan int) {
	fmt.Printf("----------writeData begin--------\n")
	for i := 0; i < 10000; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Printf("----------writeData end--------\n")
}

// 读出数据并判断是素数放进tempChan
func readData(intChan chan int, tempChan chan int, resultChan chan bool) {
	fmt.Printf("----------readData begin--------\n")
	for v := range intChan {
		if v%2 != 0 {
			tempChan <- v
		}
	}

	resultChan <- true
	fmt.Printf("----------readData.resultChan = true  end--------\n")
}
