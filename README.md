# src/sample : go 基础
go study


# go version : 1.22.1(本demo 使用的版本)
# go sdk:https://golang.google.cn/dl/
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


2. 同级文件夹下的引用并引入第三方uuid包
> 文件目录结构如下所示：
>
> chapter11_file/demo03_LogAnalysis/
> +  LogEntity.go
> + main.go

2.1 引入第三方uuid包 在文件夹 `chapter11_file/` 执行 `go mod init demo03_LogAnalysis` 后再执行 `go mod tidy` 生成`go.mod`文件

2.2 在执行命令 `go get github.com/google/uuid`

2.3  在 `main.go` 中引用`LogEntity.go`
> 创建 `LogEntity.go` 文件声 明相同的包名 `package main` ，Go 编译器会自动处理它们之间的依赖关系

2.4 如果执行报错 undefined
> Go 中 main 包默认不会加载其他文件， 而其他包都是默认加载的。如果 main 包有多个文件，则在执行的时候需要将其它文件都带上，即执行 go run *.go
> (可能跟安装的sdk版本不一致有关：1.22.1不会出现这个问题 1.20.5就会有问题)


3. 安装 redis 
> PS D:\Y\projects\personal\imtuddou.go\src\sample\chapter_17_redis> go mod init demo01

> PS D:\Y\projects\personal\imtuddou.go\src\sample\chapter_17_redis> go mod tidy

> PS D:\Y\projects\personal\imtuddou.go\src\sample\chapter_17_redis> go get github.com/go-redis/redis


4. 测试


# src/gowebsample : go web, gin


+ https://github.com/Q1mi/go_web
