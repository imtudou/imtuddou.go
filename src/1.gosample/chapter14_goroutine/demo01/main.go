package main

import (
	"fmt"
	"time"
)

/*
go 协程和主线程
Go主线程(有程序员直接称为线程/也可以理解成进程):一个Go线程上,可以起多个协程,你可以这样理解,协程是轻量级的线程[编译器做优化]。
*/

// 编写一个函数 每隔1s输出 hello world

func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d,test() hello world\n", i)
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启一个协程

	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d,main() hello world\n", i)
		time.Sleep(time.Second)
	}

	/*
		1)主线程是一个物理线程,直接作用在cpu上的。是重量级的,非常耗费cpu资源。
		2)协程从主线程开启的,是轻量级的线程,是逻辑态。对资源消耗相对小。
		3) Golang的协程机制是重要的特点,可以轻松的开启上万个协程。其它编程语言的并发机制是一般基于线程的,开启过多的线程,资源耗费大,这里就突显Golang在并发上的优势了*/
}
