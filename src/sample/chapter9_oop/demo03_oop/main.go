package main

import (
	"fmt"
	account "sample/chapter9/model/account"
)

/*
面向对象：封装
*/
func main() {
	// p := person.CreatePerson("toms")
	// p.SetAge(18)
	// p.SetSal(18000)
	// fmt.Println("person=", *p, "name=", p.Name, "age=", p.GetAge(), "sal=", p.GetSal())
	// fmt.Println("==============================================")

	var no string
	var pwd string
	var flag int
	fmt.Println("请输入账号:")
	fmt.Scanln(&no)
	fmt.Println("请输入密码:")
	fmt.Scanln(&pwd)

	account := account.CreateAccount(no, pwd)
	if account == nil{
		fmt.Println("登录失败:",account.Msg)
	} else {
		fmt.Println("登录成功")
	}

	for i := 0; i < 100; i++ {
		fmt.Println("请选择您要办理的业务:1.存款 2.取款 3.查询余额 4.修改密码 5.退出")
		fmt.Scanln(&flag)
		switch flag {
		case 1:
			var addbalance float64
			fmt.Println("请输入存款金额:")
			fmt.Scanln(&addbalance)
			account.AddBalance(addbalance)

		case 2:
			var editbalance float64
			fmt.Println("请输入取款金额:")
			fmt.Scanln(&editbalance)
			account.EditBalance(editbalance)

		case 3:
			account.QueryBalance()

		case 4:
			var oldpwd string
			var newpwd string
			fmt.Println("请输入原密码:")
			fmt.Scanln(&oldpwd)
			fmt.Println("请输入新密码:")
			fmt.Scanln(&newpwd)
			account.EditPWd(oldpwd, newpwd)

		default:
			fmt.Println("选项错误！")
		}

		if flag == 5 {
			fmt.Println("退出成功")
			return
		}
		
	}
	 
}
