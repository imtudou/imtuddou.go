package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 在指定的地址和端口上监听连接
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error creating listener:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8000")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn) // 使用 goroutine 来处理连接，以便服务器可以同时处理多个连接
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// 从连接中读取数据
		text := scanner.Text()
		fmt.Println("Received:", text)

		// 发送响应回客户端
		_, err := conn.Write([]byte("Message received: " + text + "\n"))
		if err != nil {
			fmt.Println("Error writing to connection:", err.Error())
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err.Error())
	}
	fmt.Println("Client disconnected:", conn.RemoteAddr())
}
