package main

/*
面向对象：多态
2.多态数组
*/

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
	fmt.Println("手机开始工作。。。")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	Name string
}

// 让 Camera 实现 Usb接口 方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

/*
面向对象：多态
2.多态数组
*/
func main() {

	// 定义一个usb接口数组 可以存放多重结构体变量(Phone,Camera)
	var usbArr [3]Usb
	usbArr[0] = Phone{Name: "小米"}
	usbArr[1] = Phone{Name: "vivo"}
	usbArr[2] = Camera{Name: "索尼"}
	fmt.Println(usbArr)
}
