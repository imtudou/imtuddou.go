package main

import (
	"2.gowebsample/DBHelper"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

type Users struct {
	Id         int
	Name       sql.NullString
	Company    sql.NullString
	Title      sql.NullString
	Phone      sql.NullString
	Avatar     sql.NullString
	Gender     int
	Address    sql.NullString
	Email      sql.NullString
	Tel        sql.NullString
	ProvinceId sql.NullString
	Province   sql.NullString
	CityId     sql.NullString
	City       sql.NullString
	NameCard   sql.NullString
	CreteTime  sql.NullString
}

func main() {

	//AddUserInfo()

	user := &Users{
		Id: 15,
		Name: sql.NullString{
			String: "张三",
			Valid:  true,
		},
	}
	u1, err := user.GetUserInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u1)
	fmt.Println("u1--------------------------------\r\n")
	u2, err := user.GetUserInfoList()
	if err != nil {
		fmt.Println(err)
	}
	for _, users := range u2 {
		fmt.Println(users)
	}
	fmt.Println("u2--------------------------------\r\n")
	u3, err := user.AddUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u3)
	fmt.Println("u3--------------------------------\r\n")
}

// sql 操作写入
func AddUserInfo() {

	// 设置初始化密码 创建一个新的md5.Sum结构，将字符串转换为字节数组并计算MD5值
	hash := md5.New()
	io.WriteString(hash, "123456")
	hashBytes := hash.Sum(nil)
	// 将字节数组转换为十六进制字符串
	nameCard := hex.EncodeToString(hashBytes)

	sqlstr := "insert into users(Name,Gender,Email,Phone,NameCard,CreteTime)  values(?,?,?,?,?,?)"

	//args := []interface{}{"张三", "zs@163.com", 13800000001, nameCard, time.Now().Format("2006-01-02 15:04:05")}
	rows, err := DBHelper.Exec(sqlstr, "张三", 1, "zs@163.com", 13800000001, nameCard, time.Now().Format("2006-01-02 15:04:05"))
	//rows, err := DBHelper.Insert(sqlstr, u.Name, u.Gender, u.Phone, u.Email, u.Phone, u.NameCard, u.CreteTime)
	if err != nil {
		fmt.Println(err)
	}
	if rows > 0 {
		fmt.Println("新增成功")
	}
}

func (u *Users) AddUser() (bool, error) {

	u.Name = sql.NullString{
		String: "张三",
		Valid:  true,
	}

	u.Gender = 1

	u.Phone = sql.NullString{
		String: "13800000001",
		Valid:  true,
	}

	u.Email = sql.NullString{
		String: "zs@163.com",
		Valid:  true,
	}

	u.CreteTime = sql.NullString{
		String: time.Now().Format("2006-01-02 15:04:05"),
		Valid:  true,
	}

	// 设置初始化密码 创建一个新的md5.Sum结构，将字符串转换为字节数组并计算MD5值
	hash := md5.New()
	io.WriteString(hash, "123456")
	hashBytes := hash.Sum(nil)
	// 将字节数组转换为十六进制字符串
	u.NameCard = sql.NullString{
		String: hex.EncodeToString(hashBytes),
		Valid:  true,
	}

	sqlstr := "insert into users(Name,Gender,Email,Phone,NameCard,CreteTime)  values(?,?,?,?,?,?)"

	//args := []interface{}{"张三", "zs@163.com", 13800000001, nameCard, time.Now().Format("2006-01-02 15:04:05")}
	//rows, err := DBHelper.Exec(sqlstr, "张三", 1, "zs@163.com", 13800000001, nameCard, time.Now().Format("2006-01-02 15:04:05"))
	rows, err := DBHelper.Exec(sqlstr, u.Name, u.Gender, u.Email, u.Phone, u.NameCard, u.CreteTime)
	if err != nil {
		return false, err
	}
	if rows > 0 {
		fmt.Println("新增成功")
	}
	return true, nil
}

func (u *Users) GetUserInfo() (*Users, error) {
	// 注意这里传递 u.Id 而不是 &u.Id，除非你的 DBHelper.QueryRow 需要一个指针
	row, err := DBHelper.QueryRow("select * from users where Id=?", u.Id)
	if err != nil {
		return nil, err
	}

	user := &Users{}
	col := []interface{}{ // 上面查的* 就必须列出全部字段
		&user.Id,
		&user.Name,
		&user.Company,
		&user.Title,
		&user.Phone,
		&user.Avatar,
		&user.Gender,
		&user.Address,
		&user.Email,
		&user.Tel,
		&user.ProvinceId,
		&user.Province,
		&user.CityId,
		&user.City,
		&user.Name,
		&user.CreteTime,
	}
	err = row.Scan(col...) // 使用 ... 来展开切片
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *Users) GetUserInfoList() ([]*Users, error) {
	rows, err := DBHelper.Query("select * from users where name=?", u.Name)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// 定义返回值 user 切片
	var userList []*Users
	// 遍历
	for rows.Next() {
		user := &Users{}
		col := []interface{}{ // 上面查的* 就必须列出全部字段
			&user.Id,
			&user.Name,
			&user.Company,
			&user.Title,
			&user.Phone,
			&user.Avatar,
			&user.Gender,
			&user.Address,
			&user.Email,
			&user.Tel,
			&user.ProvinceId,
			&user.Province,
			&user.CityId,
			&user.City,
			&user.Name,
			&user.CreteTime,
		}
		err = rows.Scan(col...) // 使用 ... 来展开切片

		if err != nil {
			return nil, err
		}
		userList = append(userList, user) // 添加到切片数组中
	}
	return userList, nil
}
