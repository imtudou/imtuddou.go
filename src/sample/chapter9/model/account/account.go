package account

import "fmt"

var defaultNo string = "666666"
var defaultPwd string = "zs666666"
var defaultbalance float64 = 1000.5

type account struct {
	account_no      string
	account_pwd     string
	account_balance float64
}

// 登录
func CreateAccount(no string, pwd string) *account {
	if len(no) != 6 {
		fmt.Println("账号不正确")
	}

	if no != defaultNo {
		fmt.Println("账号不正确")
	}

	if len(pwd) < 6 {
		fmt.Println("密码不正确")
	}

	if pwd != defaultPwd {
		fmt.Println("密码不正确")
	}

	return &account{
		account_no:      no,
		account_pwd:     pwd,
		account_balance: defaultbalance,
	}
}

// 存款
func (account *account) AddBalance(balance float64) {
	if balance <= 0 {
		fmt.Println("输入金额有误!")
	}
	account.account_balance += balance
	fmt.Println("存款成功!")
}

// 取款
func (account *account) EditBalance(balance float64) {
	if balance <= 0 {
		fmt.Println("输入金额有误!")
	}

	if account.account_balance <= 0 {
		fmt.Println("余额不足!")
	}

	account.account_balance -= balance
	fmt.Println("取款成功!")
}

// 查询余额
func (account *account) QueryBalance() {
	fmt.Printf("账号:%v,余额:%v \n", account.account_no, account.account_balance)
}

// 修改密码
func (account *account) EditPWd(oldPwd string, newPwd string) {
	if len(oldPwd) < 6 {
		fmt.Println("密码不正确!")
		return
	}

	if oldPwd != defaultPwd {
		fmt.Println("密码不正确!")
	}

	if len(newPwd) < 6 {
		fmt.Println("新密码不能少于六位!")
	}
	account.account_pwd = newPwd
	fmt.Printf("修改成功，账号:%v,密码:%v \n", account.account_no, account.account_pwd)
}
