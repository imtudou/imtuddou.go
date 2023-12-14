package main

import (
	"fmt"
)

func main() {

	names := [...]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	heroName := ""
	fmt.Println("请输入名字：")
	fmt.Scanln(&heroName)

	// for i := 0; i < len(names); i++ {
	// 	if heroName == names[i] {
	// 		fmt.Printf("名字=%v,下标=%v", names[i], i)
	// 	} else if i == len(names)-1 {
	// 		fmt.Println("没有找到")
	// 	}

	// }

	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("名字=%v,下标=%v", names[i], i)
			break
		} else if i == len(names)-1 {
			fmt.Println("没有找到")
		}
	}

	if index != -1 {
		fmt.Printf("名字=%v,下标=%v", names[index], index)
	} else {
		fmt.Println("没有找到:", heroName)
	}

}
