package main

import (
	"fmt"
	"time"
)

func main() {

	// 统计test() 方法的执行时间
	begin := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("test() 方法耗时：%vs \n", end-begin)

}

func test() {
	now := time.Now()
	fmt.Printf("now=%v, type=%T \n", now, now)

	//24小时格式化输出
	fmt.Printf("%d-%d-%d %d:%d:%d \n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	// 解析时间
	// 通常我们会有一个需求就是，将一个字符串时间按照一定格式转换为Go中的时间结构体，接下来我们要做的就是这件事。
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	inlocation, err := time.ParseInLocation("2006-01-02", "2023-11-30", location)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inlocation.String())

	//time.Sleep()可以使用当前goroutine处于挂起状态一定的时间，在这期间goroutine将被阻塞，直到恢复运行状态。
	for i := 0; i < 5; i++ {
		fmt.Println(now.Format("2006/01/02 15:04:05"))
		time.Sleep(time.Second * 2) // 2s
	}

	// Timer
	// Timer是一个计时器，对外暴露一个channel，当指定时间到了以后，channel就会收到消息并关闭。
	//time.NewTimer(time.Second)

	// 	Ticker
	// Ticker是一个定时器，与timer的区别在于，timer是一次性的，而Ticker是定时触发。
	//通过time.NewTicker()可以创建一个新的定时器

	// time的Unix 和 UnixNano
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}
