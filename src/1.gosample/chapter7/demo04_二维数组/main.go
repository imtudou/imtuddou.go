package main

import "fmt"

func main() {
	var arr [2][3]int
	arr[1][1] = 10
	fmt.Println(arr)

	fmt.Printf("arr[0]的地址是%p\n", &arr[0])
	fmt.Printf("arr[1]的地址是%p\n", &arr[1])

	fmt.Printf("arr[0][0]的地址是%p\n", &arr[0][0])
	fmt.Printf("arr[1][0]的地址是%p\n", &arr[1][0])

	// 演示二维数组的遍历
	var arr1 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Println(arr1[i][j])
		}
	}

	fmt.Println()

	for i, v := range arr1 {
		for j, jv := range v {
			fmt.Printf("arr1[%v][%v] = %v \n", i, j, jv)
		}
	}
	fmt.Println()

	var studentScores [3][5]float64
	for i := 0; i < len(studentScores); i++ {
		for j := 0; j < len(studentScores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩 \n", i+1, j+1)
			fmt.Scanln(&studentScores[i][j])
		}
	}

	for i := 0; i < len(studentScores); i++ {
		sum := 0.0
		for j := 0; j < len(studentScores[i]); j++ {
			sum += studentScores[i][j]
		}
		fmt.Printf("第%v班的总分是:%v,平均分是:%v \n", i, sum, sum/float64(len(studentScores[i])))
	}

}
