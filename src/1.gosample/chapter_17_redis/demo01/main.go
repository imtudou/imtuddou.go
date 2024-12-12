package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {

	// Context 是 Go 1.7 版本引入的一个标准库接口，用于在不同的 Goroutine 之间传递截止日期、取消信号以及其他请求范围的值、取消信号、超时时间等。它非常适用于控制并发操作的生命周期，如 HTTP 请求、数据库查询等
	//context := context2.Background()

	//1. 连接redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // 如果Redis服务器设置了密码，请在这里填写
		DB:       0,                // 使用哪个数据库（Redis默认有16个数据库，编号从0到15
	})

	// 1. 字符串操作: 设置键值对
	client.Set("key", "value", 0).Err()

	// 获取
	val, _ := client.Get("key").Result()
	fmt.Println(val)

	// 删除key
	client.Del("key")

	b := SetString(client, "cs", "cscscscsc")
	if b {
		fmt.Println("写入redis成功")
	} else {
		fmt.Println("写入redis失败")
	}

}

func SetString(c *redis.Client, key string, value string) bool {

	ret := c.Set(key, value, time.Second*20) // 20s
	if ret.Err() != nil {
		return false
	}
	return true
}
