package main

import (
	"fmt"
)

// 申明Usb接口
type Usb interface {
	// 定义两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让Phone 实现 Usb接口 方法
func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作。。。")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	Name string
}

// 让 Camera 实现 Usb接口 方法
func (c Camera) Start() {
	fmt.Println(c.Name, "相机开始工作。。。")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type T interface {
}

/*
面向对象：多态
1.多态参数
*/
func main() {

	//	var p Phone // 结构体变量 实现了Usb.Start()
	var usb Usb = Phone{Name: "小米"}
	usb.Start()
	usb.Stop()

	var t1 T = Camera{Name: "索尼"}
	fmt.Println(t1)

	var t2 interface{} // 空接口可以接收任何值
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2)
}
