package utils

import (
	"fmt"
	"time"
)

var HeroName string = "吕布"
var Age int
var Name string

func init() {
	Age = 111
	Name = "张飞"
	fmt.Println("utils.init()...")

}

func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64

	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
	}
	return res
}

func Now() time.Time {
	return time.Now()
}
