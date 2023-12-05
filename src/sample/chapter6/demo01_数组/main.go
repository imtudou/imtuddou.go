package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1.数组的地址可以通过数组名来获取 &nums
2.数组的第一个元素的地址就是数组的首地址
3.数组的各个元素的地址间隔是依据数组的类型决定的 比如int64  +8
int32 +4
*/
func main() {

	fmt.Println("--------------------------")
	var nums1 [5]int         // 初始化
	nums2 := [5]int{1, 2, 3} // 也可以用元素初始化
	//nums3 := new([5]int)     // 还可以通过new函数获得一个指针

	//这里的[...]是规定写法
	var nums4 = [...]int{11, 22, 33}
	var nums5 = [...]int{1: 11, 0: 22, 2: 33, 4: 666}

	fmt.Println(nums1[0])
	fmt.Println(cap(nums1))
	fmt.Println(len(nums2)) // 还可以通过内置函数len来访问数组元素的数量
	//fmt.Println(cap(nums3)) // 内置函数cap来访问数组容量，数组的容量等于数组长度，容量对于切片才有意义
	fmt.Println(nums4)
	fmt.Println(nums5)
	fmt.Println("--------------------------")
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
	heros := [...]string{"zs", "ls", "ww", "mz"}
	for i, v := range heros {
		fmt.Printf("i=%v,v =%v\n", i, v)
		fmt.Printf("heros[%d]=%v\n", i, heros[i])
	}
	fmt.Println("--------------------------")
	//1）创建一个byte类型的26个元素的数组，分别 放置'A'-'Z'。使用for循环访问所有元素并打印出来。提示：字符数据运算'A'+1->'B'
	byteArr := [26]byte{}
	for i := 0; i < cap(byteArr); i++ {
		byteArr[i] = 'A' + byte(i)
		fmt.Printf("%c ", byteArr[i])
	}

	//2)请求出一个数组的最大值,并得到对应的下标。
	intArr := [...]int{1, 23, 456, -109, 0, 90}
	maxVal := intArr[0]
	maxIndex := 0

	for i := 0; i < len(intArr); i++ {
		if intArr[i] > maxVal {
			maxVal = intArr[i]
			maxIndex = i
		}
	}
	fmt.Println("maxVal:", maxVal)
	fmt.Println("maxIndex:", maxIndex)
	fmt.Println("--------------------------")
	//3)请求出一个数组的和和平均值。for-range
	sum := 0
	for _, v := range intArr {
		sum += v
	}
	fmt.Println("sum=", sum)
	fmt.Println("avg=", sum/len(intArr))
	fmt.Println("avg(保留小数)=", float64(sum)/float64(len(intArr)))
	fmt.Println("--------------------------")
	//4.要求：随机生成五个数，并将其反转打印 指定范围随机数（0~100）
	intArr1 := [5]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr1); i++ {
		intArr1[i] = rand.Intn(100)
	}
	fmt.Println("intArr1=", intArr1)
	// 反转打印 交换的次数是len(intArr1)/2
	temp := 0
	for i := 0; i < len(intArr1)/2; i++ {
		temp = intArr1[len(intArr1)-1-i]
		intArr1[len(intArr1)-1-i] = intArr1[i]
		intArr1[i] = temp

	}
	fmt.Println("反转后intArr1=", intArr1)

}
