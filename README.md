# imtuddou.go
go study

#  官方文档 http://doc.golang.ltd/


# chapter3 Golang 变量
#### 为什么需要变量
#### 变量的介绍
#### 变量快速入门
#### 变量使用注意事项
#### 变量申明，初始化和赋值
#### 程序中+号的使用
#### 数据类型基本介绍
#### 整数
#### 浮点数
字符类型本质探讨
1)字符型存储到计算机中,需要将字符对应的码值(整数)找出来
    存储：字符--->对应码值---->二进制-->存储
    读取：二进制---->码值--->字符-->读取
    
2)字符和码值的对应关系是通过字符编码表决定的(是规定好)
Go语言的编码都统一成了utf-8,和其它的编程语言来说。非常的方便，很统一，在也有编码的困扰了


# chapter5 Golang 包引用
>示例：
现在需要在`src\sample\chapter5\demo01\main.go` 文件中导入 `src\sample\chapter5\utils\utils.go`问文件中的 `Cal()`方法 则需要在 `src\sample\chapter5` 文件中执行下面的命令

1. 解决  `package xxx/xxx/xxx is not in GOROOT (xxx)` 问题
> 1. 开启mod 模式 `go env -w GO111MODULE=on`
> 2. `go mod init` 在当前项目文件下生成 go.mod文件 
>3. 接着执行`go mod tidy`

