package main

import (
	"fmt"
)

func main() {
	//第一种使用方式
	var a map[string]string
	a = make(map[string]string)
	a["key1"] = "value1"
	a["key2"] = "value2"
	fmt.Println(a)

	//第二种使用方式
	mp := make(map[string]string)
	mp["key1"] = "value1"
	mp["key2"] = "value2"
	mp["key1"] = "value2"
	fmt.Println(mp)

	//第三种使用方式
	heros := map[string]string{
		"hero1": "松江",
		"hero2": "吴用",
	}
	heros["hero3"] = "林冲"
	fmt.Println(heros)

	//案例/*课堂练习：演示一个key-value 的value是map的案例比如：我们要存放3个学生信息，每个学生有 name和sex 信息思路：map[string]map[string]string

	stuMap := make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string)
	stuMap["stu01"]["name"] = "张三"
	stuMap["stu01"]["sex"] = "男"

	stuMap["stu02"] = make(map[string]string)
	stuMap["stu02"]["name"] = "小红"
	stuMap["stu02"]["sex"] = "女"
	fmt.Println(stuMap)

	fmt.Println("==================")
	for key, val := range stuMap {
		fmt.Printf("key=%v,value=%v \n", key, val)
		for subkey, subVal := range val {
			fmt.Printf("subkey=%v,subVal=%v \n", subkey, subVal)
		}

	}
	fmt.Println("==================")
	// 查找
	val, ok := stuMap["stu02"]
	if ok {
		fmt.Printf("ok=%v,value=%v \n", ok, val)

	}

	// 删除 stu01 的数据
	delete(stuMap, "stu01")
	fmt.Println(stuMap)

	// 一次性删除所有的key
	// 1.遍历所有的key逐一删除
	// 2.直接在重新赋值 stuMap = make(map[string]string) 或者  stuMap = nil
	stuMap = nil
	fmt.Println(stuMap)

}
