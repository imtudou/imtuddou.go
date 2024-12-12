package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string `json:"name"` // struct tag 结构体标签
	Age   int    `json:"age"`
	Color string `json:"color"`
	Hobby string `json:"hoby"`
}

func (c Cat) getName() {
	fmt.Println("getName()", c.Name)
}

func main() {
	// 张老太养了20只猫猫：一只名字叫小白，今年3岁，白色。还有一只叫小花，今年100岁，花色。请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。如果用户输入的小猫名错误，则显示 张老太没有这只猫猫。

	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"
	fmt.Println(cat1)

	fmt.Println("输出如下：")
	fmt.Println("Name:", cat1.Name)
	fmt.Println("Age:", cat1.Age)
	fmt.Println("Color:", cat1.Color)
	fmt.Println("Hobby:", cat1.Hobby)

	// 方式二
	cat2 := Cat{"zs", 111, "xxxxx", "xxxxx"}
	fmt.Println(cat2)
	cat2.getName()

	// json 序列化

	jsonstr, err := json.Marshal(cat2)
	if err != nil {
		fmt.Println("json 序列化错误！")
	}
	fmt.Println("cat2 json:", string(jsonstr))

}
