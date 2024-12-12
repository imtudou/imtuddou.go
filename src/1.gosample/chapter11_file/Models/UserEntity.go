package Models

import "fmt"

type UserEntity struct {
	Name   string `json:"name"`   //姓名
	Age    int    `json:"age"`    // 年龄
	Gender int    `json:"gender"` // 性别 0:男 1：女
}

func GetUser() string {

	user := UserEntity{Name: "zs", Age: 10, Gender: 1}
	return fmt.Sprintf("调用GetUser Name = %v,Age = %v,Gender = %v", user.Name, user.Age, user.Gender)
}
