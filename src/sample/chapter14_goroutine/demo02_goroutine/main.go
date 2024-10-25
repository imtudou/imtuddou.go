package main

import (
	"fmt"
	"sync"
)

var (
	mapList = make(map[int]int)

	/*
		申明一个全局互斥锁
		lock: 是一个全局互斥锁
		sync 是包: synchornized 同步
		Mutex: 是互斥
	*/
	lock sync.Mutex
	wg   sync.WaitGroup
)

/*
//需求:现在要计算1-200的各个数的阶乘,并且把各个数的阶乘放入到map 中。//最后显示出来。要求使用goroutine完成//

思路
//1,编写一个函数,来计算各个数的阶乘,并放入到map中.
//2.我们启动的协程多个,统计的将结果放入到map中
*/
func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	wg.Wait()
	fmt.Printf("-----------输出结果---------------- \n")

	for i, i2 := range mapList {
		fmt.Printf("map[%v] = %v \n", i, i2)
	}
}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock() // 加锁
	mapList[n] = res
	lock.Unlock() // 释放锁
}
