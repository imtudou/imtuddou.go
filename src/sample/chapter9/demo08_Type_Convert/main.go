package main

import "fmt"

type Student struct {
}

// 写一个函数 循环判断参数类型

func TypeConvert(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("value=%v,type=%T \n", v, v)
		case float64:
			fmt.Printf("value=%v,type=%T \n", v, v)
		case int, int64:
			fmt.Printf("value=%v,type=%T \n", v, v)
		case string:
			fmt.Printf("value=%v,type=%T \n", v, v)
		case nil:
			fmt.Printf("value=%v,type=%T \n", v, v)
		case Student:
			fmt.Printf("value=%v,type=%T \n", v, v)
		case *Student:
			fmt.Printf("value=%v,type=%T \n", v, v)
		default:
			fmt.Println("无匹配类型")
		}
	}

}

/*
类型转换
*/

func main() {
	p1 := true
	p2 := 9.98
	p3 := 102533
	p4 := "222"
	p5 := "222"
	p6 := Student{}
	p7 := &Student{}

	TypeConvert(p1, p2, p3, p4, p5, p6, p7)

}
