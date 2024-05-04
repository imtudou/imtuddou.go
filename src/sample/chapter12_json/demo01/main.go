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

	// 序列化
	str, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("序列化 %v\n", string(str))

	// 1.反序列化成 struct
	//unstr := "{\"monster_name\":\"张三\",\"monster_age\":1111,\"monster_Birthday\":\"2012-22-21\",\"monster_Sal\":850.045,\"monster_Skill\":\"你叉叉\"}"
	var unmonster Monster
	un_struct_err := json.Unmarshal([]byte(str), &unmonster)
	if un_struct_err != nil {
		fmt.Println(un_struct_err)
	}
	fmt.Println(unmonster)

	// 2.反序列化成 map
	var unmap map[string]interface{}
	un_map_err := json.Unmarshal([]byte(str), &unmap)
	if un_map_err != nil {
		fmt.Println(un_map_err)
	}
	fmt.Println(unmap)

	// 3.反序列化成 Slice
	unslicestr := "[{\"monster_name\":\"张三\",\"monster_age\":1111,\"monster_Birthday\":\"2012-22-21\",\"monster_Sal\":850.045,\"monster_Skill\":\"你叉叉\"}]"
	var unslice []map[string]interface{}
	un_slice_err := json.Unmarshal([]byte(unslicestr), &unslice)
	if un_slice_err != nil {
		fmt.Println(un_slice_err)
	}
	fmt.Println(unslice)

	fmt.Println("==================================")
	// xml 数据格式

	monsterxml := "<Monster><Name>张三</Name><Age>1111</Age><Skill>你叉叉</Skill></Monster>"
	xmlReder := strings.NewReader(monsterxml)
	p := xml.NewDecoder(xmlReder)
	fmt.Printf("p=%v\n", p)
}
