package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

/*
		编写一个服务端程序 监听 8000 端口
	    1. 可以和多个客户端常见连接
	 	2. 连接成功后 客户端发送数据 服务端接受数据并显示
*/
func main() {

	// 1. tcp表示网络协议是tcp
	// 2. 监听本地8080端口
	lister, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("lister err =", err)
		return
	}
	defer lister.Close() // 延时关闭lister
	fmt.Println("服务开始监听8000......", time.Now().Format("2006-01-02 15:04:05"))

	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端连接：", time.Now().Format("2006-01-02 15:04:05"))
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("lister err =", err)
			continue
		}

		// 这里准备一个协程 为客户端服务
		go handlerConn(conn)
	}

}

func handlerConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("客户端已连接:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	// 循环接受客户端发送的数据
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
}
