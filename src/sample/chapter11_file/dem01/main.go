package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("../../../../doc/111.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	// 函数退出时关闭文件
	defer file.Close()

	// 使用*reader 对象读取 带缓冲区
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 表示文件结尾
			break
		}
		fmt.Println(str)
	}

}
