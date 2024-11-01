package main

import (
	"fmt"
	"reflect"
)

func main() {

	//1)使用反射来遍历结构体的字段,调用结构体的方法,并获取结构体标签的值
	m := Monster{Name: "zs", Age: 200, Gender: "男"}
	//reflectTest(m)
	fmt.Println("------------------------------分割线------------------------------ ")
	reflectTestPtr(&m)
}

func reflectTest(obj interface{}) {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	// 获取到 obj 对应的类别
	k := v.Kind()
	if k != reflect.Struct {
		// 如果传入的不是 struct 就退出
		fmt.Println("传入的类型错误 = ", t)
		return
	}

	// 获取到该结构体有几个字段并遍历
	length := t.NumField()
	for i := 0; i < length; i++ {
		field := t.Field(i).Name
		value := v.FieldByName(field).Interface()
		tag := t.Field(i).Tag.Get("json") // 获取tag 标签的值
		// 如果该字段是tag标签就显示
		if tag != "" {
			fmt.Printf("遍历结构体%v：filed = %v, value = %v , tag = %v \n", i, field, value, tag)
		} else {
			fmt.Printf("遍历结构体%v：filed = %v, value = %v \n", i, field, value)
		}

	}

	// 获取结构体所有方法
	numMethod := t.NumMethod()
	fmt.Printf("NumMethod= %v\n", numMethod)

	// 反射调用方法(有参有返回值)
	params1 := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	ret1 := v.MethodByName("GetAge").Call(params1)
	if len(ret1) > 0 {
		// 处理返回值
		if ret1[0].IsValid() {
			fmt.Println("调用 GetAge方法成功：ret1[0] =", ret1[0])
		}

		if len(ret1) > 1 && ret1[1].IsValid() {
			fmt.Println("调用 GetAge方法成功：ret1[1] =", ret1[1])
		}
	}

	// 如果您直接传递了 Monster 实例的反射值（而不是指针），则 Set 方法内的修改将不会反映到原始实例上 详见 SetPtr()
	params2 := []reflect.Value{reflect.ValueOf("小三"), reflect.ValueOf(400), reflect.ValueOf("女")}
	ret2 := v.MethodByName("Set").Call(params2)
	fmt.Println("调用 Set 方法成功,不会修改值", ret2)
	v.MethodByName("Print").Call(nil)

}

// 测试反射调用指针方法
func reflectTestPtr(obj interface{}) {
	v := reflect.ValueOf(obj) // 获取指针实际值
	fmt.Println("v = ", v)
	params2 := []reflect.Value{reflect.ValueOf("小三"), reflect.ValueOf(239), reflect.ValueOf("女")}
	fmt.Println("params2 = ", params2)
	ret2 := v.MethodByName("SetPtr").Call(params2)
	fmt.Println("调用 SetPtr 方法成功,会修改值", ret2)
	v.MethodByName("Print").Call(nil)

}

// 给Monster 绑定的方法
func (s Monster) GetAge(a1 int, a2 int) int {
	fmt.Println("调用 GetAge 方法 --------- ")
	return a1 + a2
}

// 给Monster 绑定的方法
func (s Monster) Set(name string, age int, gender string) {
	fmt.Println("调用 Set 方法 --------- ")
	s.Name = name
	s.Age = age
	s.Gender = gender
}

// 给Monster 绑定的方法
func (s *Monster) SetPtr(name string, age int, gender string) {
	fmt.Println("调用 SetPtr 方法 --------- ")
	s.Name = name
	s.Age = age
	s.Gender = gender
}

func (s Monster) Print() {
	fmt.Println("调用 Print 方法验证值 --------- ")
	fmt.Println(s)
}

type Monster struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
