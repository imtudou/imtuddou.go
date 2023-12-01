package main

import "fmt"

/*
1.数组的地址可以通过数组名来获取 &nums
2.数组的第一个元素的地址就是数组的首地址
3.数组的各个元素的地址间隔是依据数组的类型决定的 比如int64  +8
int32 +4
*/
func main() {

	var nums1 [5]int         // 初始化
	nums2 := [5]int{1, 2, 3} // 也可以用元素初始化
	nums3 := new([5]int)     // 还可以通过new函数获得一个指针
	//这里的[...]是规定写法
	var nums4 = [...]int{11, 22, 33}
	var nums5 = [...]int{1: 11, 0: 22, 2: 33}

	fmt.Println(nums1[0])
	fmt.Println(len(nums2)) // 还可以通过内置函数len来访问数组元素的数量
	fmt.Println(cap(nums3)) // 内置函数cap来访问数组容量，数组的容量等于数组长度，容量对于切片才有意义
	fmt.Println(nums4)
	fmt.Println(nums5)

	// 从终端输入5个成绩保存保存到float64数组中
	// source := [5]float64{}
	// for i := 0; i < len(source); i++ {
	// 	fmt.Printf("请输入%d个元素的值\n", i)
	// 	fmt.Scanln(&source[i])
	// }
	// for i := 0; i < len(source); i++ {
	// 	fmt.Printf("source[%d]=%v\n", i, source[i])
	// }

	//数组遍历 for-range go语言一种独特的结构遍历

}
