package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// 从标准输入读取用户输入，并发送到服务器
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')

		// 发送数据到服务器
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to server:", err.Error())
			break
		}

		// 从服务器读取响应
		scanner := bufio.NewScanner(conn)
		if scanner.Scan() {
			response := scanner.Text()
			fmt.Println("Server response:", response)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from server:", err.Error())
			break
		}
	}
}
