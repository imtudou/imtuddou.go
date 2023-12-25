package main

import (
	"errors"
	"fmt"
)

// 使用defer + recover捕获和处理异常
func Calc(n1 int, n2 int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("发送邮件......")
		}
	}()
	res := n1 / n2
	return res
}

// 自定义错误 	error.New(type)
func Calc2(n1 int, n2 int) (res int, err error) {
	defer func() {
		errres := recover()
		if errres != nil {
			// panic(err)
			err = errors.New("error:计算异常")
		}
	}()
	res = n1 / n2
	return res, nil
}

// 自定义错误 	error.New(type)
func Calc3(n1 int, n2 int) (res int, err error) {
	if n2 <= 0 {
		return 0, errors.New("分母必须大于0!")
	}
	res = n1 / n2
	return res, nil
}

func main() {
	// res := Calc(100, 0)
	// fmt.Println(res)

	// res, err := Calc2(100, 0)
	// fmt.Printf("res=%v, err=%v", res, err)

	res, err := Calc3(100, 0)
	if err != nil {
		// 如果异常错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Printf("res=%v, err=%v", res, err)
}
