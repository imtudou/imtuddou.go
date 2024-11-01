package main

import (
	"fmt"
	"reflect"
)

func main() {

	//reflecTest01(100)

	//stu := Student{Name: "zs"}
	//reflectTest02(stu)

	reflectTest03()
}

// 请编写一个案例,演示对(基本数据类型、interface()、 reflect.Value)进行反射的基本操作代码演示，
func reflecTest01(obj interface{}) {

	//通过反射获取传入变量的type kind 值
	//	1. 先获取 type
	t := reflect.TypeOf(obj)
	fmt.Println("type = ", t)
	fmt.Println("kind = ", t.Kind())

	// 2. 在获取value
	v := reflect.ValueOf(obj)
	fmt.Println("value = ", v)

	n1 := 5 + v.Int()
	fmt.Println("n1 value = ", n1)

	//  转成 interface{}
	ii := v.Interface()
	fmt.Printf("value = %v type = %T \n", ii, ii)

	// 使用断言校验类型
	num2 := ii.(int)
	fmt.Printf("num2 = %v type = %T \n", num2, num2)
}

// 请编写一个案例,演示对(结构体类型、interface()、 reflect.Value)进行反射的基本操作代码演示：
func reflectTest02(obj interface{}) {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Println("type = ", t) // 类型
	fmt.Println("t.kind = ", t.Kind())
	fmt.Println("v.kind = ", v.Kind()) // 类别 struct
	fmt.Println("value = ", v)

	//  转成 interface{}
	ii := v.Interface()

	// 使用断言校验类型
	stu, ok := ii.(Student)
	if ok {
		fmt.Printf("stu.Name = %v type = %T \n", stu.Name, stu)
	}
}

// 使用reflect.Value.Elem() 方法修改反射变量值
func reflectTest03() {

	num := 100
	v := reflect.ValueOf(&num)
	v.Elem().SetInt(500)
	fmt.Printf("修改变量的值 = %v , type = %T \n", num, num)

	str := "zs"
	fv := reflect.ValueOf(&str)
	fv.Elem().SetString("LS")
	fmt.Printf("修改变量的值 = %v , type = %T \n", str, str)
}

type Student struct {
	Name string
	Age  int
}
