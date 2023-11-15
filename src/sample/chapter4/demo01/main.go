package main

import (
	"fmt"
)

// Go 语言运算符

func main() {

	// 算术运算符
	// arithmeticOperator()
	// 关系运算符
	// relationalOperator()
	// 逻辑运算符
	// logicalOperator()
	// 位运算符
	// 赋值运算符
	// assignmentOperator()
	// 其他运算符
	// otherOperator()

	// 演示goto 使用
	gotosample()

}

func arithmeticOperator() {
	fmt.Println(10 / 4)   // output:2
	fmt.Println(10.0 / 4) // output:2.5

	// 需要转成 float计算返回小数位
	var n1 float32 = 10.0
	fmt.Println(n1 / 4) // output:2.5

	// 演示 % 取模运算
	// a%b = a-a/b*b
	fmt.Println("10 / 3=", 10/3)
	fmt.Println("-10 / 3=", -10/3)
	fmt.Println("10 / -3=", 10/-3)
	fmt.Println("-10 / -3=", -10/-3)

	// ++ 和 -- 的使用
	// 只能单独使用

	// fmt.Println("n2 = ", n2++) // error
	// fmt.Println("n2 = ", n2-- // error

	var n2 int = 10
	// n2++
	// fmt.Println("n2 = ", n2)
	// n2--
	// fmt.Println("n2 = ", n2)

	n2 += 1
	fmt.Println("n2 = ", n2)
}

func relationalOperator() {

	var a int = 21
	var b int = 10

	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

func logicalOperator() {
	var a bool = true
	var b bool = false

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!(a && b))
}

func assignmentOperator() {

	var a int = 21
	var c int
	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)
}

func otherOperator() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

func gotosample() {
	var a int = 11

	fmt.Println("ok1")
	if a >= 10 {
		goto lable1
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
lable1:
	fmt.Println("ok4")
	fmt.Println("ok5")
	// output: ok1 ok4 ok5
}
