package main

import "fmt"

func bubb(arr *[]int) {
	fmt.Println("排序之前arr=", (*arr))
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {

				// 交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}

		}
		fmt.Printf("第%v次排序之前arr=%v \n", i, (*arr))
	}

}

func main() {

	arr := []int{36, 12, 86, 99, 27, 13}
	bubb(&arr)
	fmt.Println("main.arr=", arr)
}

// 36, 12, 86, 99, 27, 13

// 第一轮
// 0: 12,36,  86, 99, 27, 13
// 1: 12,36,  86, 99, 27, 13
// 2: 12,36,  86, 99, 27, 13
// 3: 12,36,  86, 99, 27, 13
// 4: 12,36,  86, 27, 99, 13
// 5: 12,36,  86, 27, 13, 99
