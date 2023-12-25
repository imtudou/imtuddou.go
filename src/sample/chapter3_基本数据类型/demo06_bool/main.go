package main

import (
	"fmt"
	"unsafe"
)

// go bool类型使用
func main() {

	var flag = false
	fmt.Println("flag=", false, "占用空间:", unsafe.Sizeof(flag))
}
