package main

import (
	"demo03_LogAnalysis/Models"
	"fmt"
	"time"
)

func main() {

	begintime1 := time.Now()
	fmt.Printf("--------------------初始化日志文件开始: %v \n", begintime1.Format("2006-01-02 15:04:05"))
	Models.InitLogFile()
	endtime1 := time.Now()
	fmt.Printf("--------------------初始化日志文件结束:%v  执行耗时:%v \n\n", endtime1.Format("2006-01-02 15:04:05"), endtime1.Sub(begintime1))

	begintime2 := time.Now()
	fmt.Printf("--------------------分析日志文件开始: %v \n", begintime2.Format("2006-01-02 15:04:05"))
	Models.AnalysisFile()
	endtime2 := time.Now()
	fmt.Printf("--------------------分析日志文件结束:%v  执行耗时:%v \n\n", endtime2.Format("2006-01-02 15:04:05"), endtime2.Sub(begintime2))

	fmt.Printf(Models.GetUser())
	Models.Add()

}
