package main

import "fmt"

func main() {

	// 申明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "张三"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "李四"
		monsters[1]["age"] = "500"
	}
	fmt.Println(monsters)
	//  error: index out of range [2] with length 2
	// 需要使用到切片的append函数动态添加
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "狐狸精"
	// 	monsters[2]["age"] = "500"
	// }
	newMonster := map[string]string{
		"name": "狐狸精",
		"age":  "500",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

	// map 排序
	mapSort := make(map[int]int)
	mapSort[19] = 19
	mapSort[1] = 1
	mapSort[15] = 15
	mapSort[11] = 11

	fmt.Println(mapSort)
}
