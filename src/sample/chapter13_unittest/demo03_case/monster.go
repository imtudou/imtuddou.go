package demo03_case

import (
	"encoding/json"
	"errors"
	"os"
)

type Monster struct {
	Name   string
	Age    int
	Gender string
}

// 序列化并存储到本地文件
func (this *Monster) MarshalAndWriteFile() error {
	data, err := json.Marshal(this)
	if err != nil {
		return errors.New(err.Error())
	}
	err = os.WriteFile("./monster.json", data, os.ModePerm)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// 反序列化返回Nonster 对象
func (this *Monster) UnmarshalAndReadFile() error {
	data, err := os.ReadFile("./monster.json")
	if err != nil {
		return errors.New(err.Error())
	}
	json.Unmarshal(data, this)
	return nil
}
