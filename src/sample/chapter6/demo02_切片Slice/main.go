package main

import "fmt"

/*
需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

   1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
    2. 切片的长度可以改变，因此，切片是一个可变的数组。
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

func main() {

	// 演示切片基本使用
	// 创建切片的三种各种方式
	//方式1.
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	var slice1 = arr[1:3] //取arr[1]~arr[3]的元素,不包含arr[3]：[2 3]
	fmt.Println("arr=", arr)
	fmt.Println("slice1=", slice1)
	fmt.Println("slice1 len=", len(slice1))
	fmt.Println("slice1 cap=", cap(slice1))
	fmt.Println("=====================================")
	//方式2. make(type, 0)
	var slice2 []int = make([]int, 4)
	slice2[0] = 100
	slice2[1] = 200
	fmt.Println("slice2=", slice2)
	fmt.Println("slice2 len=", len(slice2))
	fmt.Println("slice2 cap=", cap(slice2))
	fmt.Println("=====================================")
	// 方式3: 定义一个切片，直接指定具体数组，使用原理类似male方式
	var strSlice []string = []string{"tome", "jake", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice len=", len(slice2))
	fmt.Println("strSlice cap=", cap(strSlice))
	fmt.Println("=====================================")

	//切片遍历的两种方式
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%v]=%v \n", i, slice1[i])
	}
	fmt.Println()
	for i, v := range slice1 {
		fmt.Printf("i=%v,v=%v ", i, v)
	}
	fmt.Println("=====================================")

	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{11, 22, 33}
	slice3 = append(slice3, 44, 55, 66)
	fmt.Println("slice3=", slice3)
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=", slice3)

	fmt.Println("=====================================")
	//切片拷贝操作
	var slice4 []int = []int{11, 22, 33}
	var slice5 = make([]int, 5)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)
}
