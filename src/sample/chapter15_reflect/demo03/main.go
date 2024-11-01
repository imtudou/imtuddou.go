package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

// 为 Cal 类型定义指针接收者的方法
func (m Cal) GetSub(num1 int, num2 int) int {
	fmt.Println("调用GetSub方法")
	return num1 - num2
}

// 为 Cal 类型定义指针接收者的方法
func (m *Cal) GetSubPtr(num1 int, num2 int) int {
	fmt.Println("调用GetSubPtr方法")
	return num1 - num2
}

// 为 Cal 类型定义指针接收者的方法
func (m *Cal) GetSubPtrNoParams() int {
	fmt.Println("GetSubPtrNoParams")
	return m.Num1 - m.Num2
}

func reflectTest(obj interface{}) {
	fmt.Println("-----------reflectTest开始-----------")
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	length := t.NumField()
	for i := 0; i < length; i++ {
		field := t.Field(i).Name
		fmt.Printf("遍历结构体字段值: Field = %v, value = %v \n", field, val.Field(i))
	}

	num1 := int(val.Field(0).Int()) // val.Field(0).Int() 返回的是 int64 须要用使用int() 转成nt
	num2 := int(val.Field(1).Int())
	params := []reflect.Value{reflect.ValueOf(num1), reflect.ValueOf(num2)}
	fmt.Println("params=", params)
	ret := val.MethodByName("GetSub").Call(params)
	if len(ret) > 0 {
		fmt.Println("返回值=", ret[0].Interface())
	}
	fmt.Println("-----------reflectTest结束-----------")
}

func reflectTestPtr(obj interface{}) {
	fmt.Println("-----------reflectTestPtr开始-----------")
	val := reflect.ValueOf(obj).Elem() // 来获取指针指向的实际值 这里需要注意的是，如果 obj 不是指针类型，或者 obj 是 nil，那么 reflect.ValueOf(obj).Elem() 将导致 panic
	length := val.NumField()
	for i := 0; i < length; i++ {
		field := val.Field(i)
		fmt.Printf("遍历结构体字段值: Field = %v, value = %v \n", field, val.Field(i))
	}

	num1 := int(val.Field(0).Int()) // val.Field(0).Int() 返回的是 int64 须要用使用int() 转成nt
	num2 := int(val.Field(1).Int())
	params := []reflect.Value{reflect.ValueOf(num1), reflect.ValueOf(num2)}
	fmt.Println("params=", params)
	ret := val.MethodByName("GetSub").Call(params)
	if len(ret) > 0 {
		fmt.Println("返回值=", ret[0].Interface())
	}
	fmt.Println("-----------reflectTestPtr结束-----------")
}

func reflectTestPtrNoParams(obj interface{}) {
	fmt.Println("-----------reflectTestPtrNoParams开始-----------")
	val := reflect.ValueOf(obj).Elem()
	length := val.NumField()
	for i := 0; i < length; i++ {
		field := val.Field(i)
		fmt.Printf("遍历结构体字段值: Field = %v, value = %v \n", field, val.Field(i))
	}

	num1 := int(val.Field(0).Int()) // val.Field(0).Int() 返回的是 int64 须要用使用int() 转成nt
	num2 := int(val.Field(1).Int())
	params := []reflect.Value{reflect.ValueOf(num1), reflect.ValueOf(num2)}
	fmt.Println("params=", params)
	ret := val.MethodByName("GetSubPtrNoParams").Call(params)
	if len(ret) > 0 {
		fmt.Println("返回值=", ret[0].Interface())
	}
	fmt.Println("-----------reflectTestPtrNoParams结束-----------")
}

func main() {
	c := Cal{Num1: 20, Num2: 13}
	reflectTest(c)
	reflectTestPtr(&c)
	fmt.Println("Cal-----------", c)
}
