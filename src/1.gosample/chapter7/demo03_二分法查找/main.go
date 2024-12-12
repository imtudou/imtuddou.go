package main

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("未找到")
		return
	}
	middleVal := (leftIndex + rightIndex) / 2
	if (*arr)[middleVal] > findVal {
		BinaryFind(arr, leftIndex, middleVal-1, findVal)
	} else if (*arr)[middleVal] < findVal {
		BinaryFind(arr, middleVal+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到了下标：%v \n", middleVal)
	}

}

func main() {
	// 注意 二分法查找必须是一个有序数组，并且从小到大排序
	// arr := [6]int{1, 9, 210, 102, 11}
	arr := [6]int{1, 9, 11, 102, 210}
	BinaryFind(&arr, 0, len(arr)-1, 102)

}
