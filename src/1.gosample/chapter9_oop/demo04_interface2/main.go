package main

import "fmt"

// monkey 结构体
type Monkey struct {
	Name string
}

type Bird interface {
	Flying()
}

type Fish interface {
	Swimming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "会爬树。。。")
}

// LittleMonkey 结构体
type LittleMonkey struct {
	Monkey
}

// LittleMonkey实现Bird
func (littlemonkey *LittleMonkey) Flaying() {
	fmt.Println(littlemonkey.Name, "通过学习会飞翔")
}

// LittleMonkey实现Fish
func (littlemonkey *LittleMonkey) Swimming() {
	fmt.Println(littlemonkey.Name, "通过学习会游泳")
}

func main() {

	// 创建一个LittleMonkey实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flaying()
	monkey.Swimming()

}
