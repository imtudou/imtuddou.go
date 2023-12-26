package main

import (
	"fmt"
)


/*
项目1：家庭收支记账软件
*/
func main() {

	var key int
	var isExit bool = false
	var isCheckExit string
	var details string = "收支\t金额\t余额\t说明\t"
	var shouzhiStr string =""
	var money float64
	var balance float64 = 1000
	var remark string =""


	for {
		fmt.Println("\n----------------------家庭收支记账软件----------------------")
		fmt.Println("1.收支明细")
		fmt.Println("2.登记收入")
		fmt.Println("3.登记支出")
		fmt.Println("4.退出软件")
		fmt.Println("请选择:1~4")
		fmt.Scanln(&key)
		switch key {
			case 1:
				fmt.Println("【1.收支明细】")
				fmt.Println(details)
			
			case 2:
				shouzhiStr = "收入"
				fmt.Println("【2.登记收入】")
				fmt.Println(shouzhiStr,"金额:")
				fmt.Scanln(&money)
				fmt.Println(shouzhiStr,"说明:")				
				fmt.Scanln(&remark)
				balance += money
				details+= fmt.Sprintf("\n%v\t%v\t%v\t%v\t",shouzhiStr,money,balance,remark)
			case 3:
				shouzhiStr = "支出"
				fmt.Println("【3.登记支出】")	
				fmt.Println(shouzhiStr,"金额:")
				fmt.Scanln(&money)
				fmt.Println(shouzhiStr,"说明:")				
				fmt.Scanln(&remark)
				balance -= money
				details+= fmt.Sprintf("\n%v\t%v\t%v\t%v\t",shouzhiStr,money,balance,remark)

			case 4:
				fmt.Println("【4.退出软件】")
				isExit = true
			default :
				fmt.Println("输入选项有误")	
				break	
		}

		if isExit {
			fmt.Println("确定退出?(y/n)")
			fmt.Scanln(&isCheckExit)
			if isCheckExit == "y" {
				fmt.Println("退出成功")
				return
			}else if isCheckExit== "n" {
				isExit = false
			}else{
				fmt.Println("输入有误")
			}
		}
				
	}
	
}