package main

import (
	"fmt"
)

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

/*
面向对象：继承
*/
func main() {

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"海尔01", 5999}, Brand{"海尔一厂", "青岛"}}
	fmt.Println("tv=", tv)

	tv = TV{
		Goods{
			Price: 6999.99,
			Name:  "海尔02",
		},
		Brand{
			Name:    "海尔二厂",
			Address: "青岛",
		},
	}
	fmt.Println("tv=", tv)

	tv2 := TV2{&Goods{"TCL01", 19999}, &Brand{"TCL一厂", "四川"}}
	fmt.Println("tv2.Goods=", tv2.Goods, "tv2.Brand=", tv2.Brand)

	tv2 = TV2{
		&Goods{
			Price: 16999.99,
			Name:  "TCL02",
		},
		&Brand{
			Name:    "TCL二厂",
			Address: "四川",
		},
	}
	fmt.Println("tv2.Goods=", tv2.Goods, "tv2.Brand=", tv2.Brand)

}
