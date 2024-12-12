package main

import (
	"fmt"
	"sample/chapter5/utils"
)

var age = getAge()

func getAge() int {
	fmt.Print("getAge()...\n")
	return 10
}

func init() {
	fmt.Print("init()...\n")
}

// 演示全局变量和init 的执行顺序
func main() {
	fmt.Print("age=", age)
	fmt.Print("Age=", utils.Age, "Name=", utils.Name)

	/*output
	utils.init()...
	getAge()...
	init()...
	age=10Age=111Name=张飞
	*/

}
