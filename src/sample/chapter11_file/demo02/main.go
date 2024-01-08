package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type CharCount struct {
	ChCount    int // 字符
	NumCount   int // 数字
	SpaceCount int // 空格
	OtherCount int // 其他
}

func main() {

	// const (
	// 	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	// 	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	// 	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	// 	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	// 	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	// 	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	// 	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	// 	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	// )

	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	// test6()
	// test7()
	test8()
}

// 统计英文，数字，空格，和其他字符数量
func test8() {
	filePath := "../../../../doc/222.txt"

}

func test7() {
	// 将 C:/Users/DeskTop/Pictures/111.png 拷贝到 d:/111.png

	srcFile := "C:/Users/DeskTop/Pictures/111.png"
	distFile := "d:/222.png"
	_, err := copyFile(distFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("err=%v", err)
	}

}

// 拷贝文件
func copyFile(distFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	distFile, err := os.OpenFile(distFileName, os.O_WRONLY|os.O_CREATE, 000)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	writer := bufio.NewWriter(distFile)
	defer distFile.Close()

	return io.Copy(writer, reader)

}

// 判断文件是否存在
func test6() {
	filePath1 := "../../../../doc/222.txt"
	fmt.Println(FileExists(filePath1))
	filePath1 = "22245.txt"
	fmt.Println(FileExists(filePath1))
}

// 判断文件是否存在

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil { // 文件或目录不存在
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// func FileExists(path string) (bool, error) {
// 	_, err := os.Stat(path)

// 	if err != nil { // 文件或目录不存在
// 		return true, nil
// 	}
// 	if os.IsNotExist(err) {
// 		return false, nil
// 	}
// 	return false, err
// }

// 将 "../../../../doc/222.txt" 的内容读取到 333.txt
func test5() {
	filePath1 := "../../../../doc/222.txt"
	filePath2 := "../../../../doc/333.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil { // 读取错误 写日志
		return
	}

	ioutil.WriteFile(filePath2, data, 000) // 在Windwos模式下 fileModule 无效随便写
	if err != nil {                        //  写入错误 写日志
		return
	}
}

// 4)打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello，北京！"
func test4() {
	filePath := "../../../../doc/222.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err=", err)
	}

	// 即使关闭file
	defer file.Close()

	// 先读取内容并显示
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF 表示文件结尾
			break
		}
		fmt.Println(str)
	}

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("hello，北京！ %v!!!\n", time.Now())
		writer.WriteString(str)
	}

	// 因为writer是带缓存的 在调用WriteString方法时 是先写缓存 所以需要调用Flush方法将缓存的诗句写到文件中
	writer.Flush()
}

// 3)打开一个存在的文件，在原来的内容追加内容'ABC! ENGLISH!"
func test3() {

	filePath := "../../../../doc/222.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err=", err)
	}

	// 即使关闭file
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("你好，go！ ABC! ENGLISH! %v!!!\n", time.Now())
		writer.WriteString(str)
	}

	// 因为writer是带缓存的 在调用WriteString方法时 是先写缓存 所以需要调用Flush方法将缓存的诗句写到文件中
	writer.Flush()

}

// 2)打开一个存在的文件中，将原来的内容覆盖成新的内容10句"你好，go！"
func test2() {

	filePath := "../../../../doc/222.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("err=", err)
	}

	// 即使关闭file
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("你好，go！ %v!!!\n", time.Now())
		writer.WriteString(str)
	}

	// 因为writer是带缓存的 在调用WriteString方法时 是先写缓存 所以需要调用Flush方法将缓存的诗句写到文件中
	writer.Flush()

}

// 1）创建一个新文件，写入内容 5句 "hello go"
func test1() {

	filePath := "../../../../doc/222.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("err=", err)
	}

	// 即使关闭file
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("hello go %v!!!\n", time.Now())
		writer.WriteString(str)
	}

	// 因为writer是带缓存的 在调用WriteString方法时 是先写缓存 所以需要调用Flush方法将缓存的诗句写到文件中
	writer.Flush()

}
