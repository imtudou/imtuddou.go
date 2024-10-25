package main

import (
	"fmt"
	"strconv"
)

/*
1） channle 本质就是一个数据结构-队列【示意图】
2)数据是先进先出【FIFO : first in first out】
3)线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全
4) channel 有类型的，一个 string 的 channel只能存放 string 类型数据。
*/

func main() {
	fmt.Printf("---------开始--------------\n")
	//test1()
	//test2()
	//test3()
	test4()
	fmt.Printf("---------结束--------------\n")

	/*
		1) channel 中只能存放指定的数据类型
		2) channle的数据放满后,就不能再放入了
		3)如果从channel取出数据后,可以继续放入
		4)在没有使用协程的情况下,如果channel数据取完了,再取,就会报dead lock
	*/
}

// 申明一个intchan 最多可以放10个int 的读取操作
func test1() {
	// 申明
	intchan := make(chan int, 10) // 必须申明容量 后续调整
	//写入
	for i := 0; i < 10; i++ {
		intchan <- i
	}

	// 读取
	length := len(intchan) //因为通道的长度在运行时是动态的 如果知道要发送的元素数量，可以硬编码循环次数
	for i := 0; i < length; i++ {
		fmt.Printf("i = %v, intchan =%v \n", i, <-intchan)
	}
}

// 申明一个mapchan 最多可以放10个map[int]string 的读取操作
func test2() {
	var mapchan = make(chan map[int]string, 10)
	m1 := make(map[int]string, 20)
	for i := 0; i < 5; i++ {
		m1[i] = string(i) + "hello m1"
	}

	m2 := make(map[int]string, 20)
	for i := 0; i < 5; i++ {
		m2[i] = string(i) + "hello m2"
	}

	mapchan <- m1
	mapchan <- m2

	// 读取
	fmt.Printf("mapchan[m1] =%v \n", <-mapchan)
	fmt.Printf("mapchan[m2] =%v \n", <-mapchan)
}

// 申明一个studentchan 最多可以放10个sturct 的读取操作
func test3() {
	var studentchan chan Student
	studentchan = make(chan Student, 10)
	for i := 0; i < 10; i++ {
		stu := Student{
			name:  "张三" + strconv.Itoa(i),
			age:   i,
			email: "张三" + strconv.Itoa(i) + "@163.com",
		}
		if i%2 == 0 {
			stu.gender = "男"
		} else {
			stu.gender = "女"
		}
		studentchan <- stu
	}
	close(studentchan) //如果channel 没有关闭 则会出现deadlock的错误，如果已关闭则会正常遍历完成后，退出遍历
	// 读取
	for v := range studentchan {
		fmt.Printf("studentchan = %v \n", v)
	}

}

// 申明一个anychan 最多可以放10个 任意类型 的读取操作
// channel 遍历和关闭的案例演示
func test4() {
	var anychan = make(chan interface{}, 20)
	anychan <- Student{
		name:   "张三",
		age:    0,
		gender: "男",
		email:  "张三@163.com",
	}

	anychan <- "校园烧烤"
	anychan <- 2024
	anychan <- "1024"
	anychan <- "bug 大会"

	close(anychan) //如果channel 没有关闭 则会出现deadlock的错误，如果已关闭则会正常遍历完成后，退出遍历
	// 读取
	for v := range anychan {
		fmt.Printf("anychan = %v \n", v)
	}

}

type Student struct {
	name   string
	age    int
	gender string
	email  string
}
