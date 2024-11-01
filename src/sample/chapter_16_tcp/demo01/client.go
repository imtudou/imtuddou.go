package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

/*
1. 编写一个客户端程序 连接到服务端的 8000 端口
2. 客户端发送数据
3. 通过终端输入数据发送服务器端
4. 服
*/
func main() {
	fmt.Println("客户端已启动......", time.Now().Format("2006-01-02 15:04:05"))
	conn, err := net.Dial("tcp", "192.168.31.67:8000")
	if err != nil {
		fmt.Println("client err = ", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端已连接：", conn.RemoteAddr())
	// 从标准输入读取用户输入，并发送到服务器
	reader := bufio.NewReader(os.Stdin) // os.Stdin  代表输入[终端]

	line, err := reader.ReadString('\n') // 从终端读取用户输入
	if err != nil {
		fmt.Println("client read ReadString err = ", err)
	}

	_, cerr := conn.Write([]byte(line)) // 发送给服务器
	if cerr != nil {
		fmt.Println("client Write err = ", cerr)
	}

	fmt.Printf("客户端发送了:%v \n", line)

	// 从服务器读取响应
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		response := scanner.Text()
		fmt.Println("Server response:", response)
	}

}
