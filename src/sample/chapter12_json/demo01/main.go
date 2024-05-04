package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// Monster 特殊序列化可配置tag
type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_Birthday"`
	Sal      float64 `json:"monster_Sal"`
	Skill    string  `json:"monster_Skill"`
}

func main() {

	monster := Monster{
		Name:     "张三",
		Age:      1111,
		Birthday: "2012-22-21",
		Sal:      850.045,
		Skill:    "你叉叉",
	}

	str, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

	fmt.Println("==================================")
	// xml 数据格式

	monsterxml := "<Monster><Name>张三</Name><Age>1111</Age><Skill>你叉叉</Skill></Monster>"
	xmlReder := strings.NewReader(monsterxml)
	p := xml.NewDecoder(xmlReder)
	fmt.Printf("p=%v\n", p)
}
