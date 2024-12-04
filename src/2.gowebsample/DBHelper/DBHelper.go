package DBHelper

import (
	Constant2 "2.gowebsample/Common/Constant"
	"2.gowebsample/DBHelper/Prodviders"
	"database/sql"
	"errors"
)

var (
	Db  *sql.DB // 全局变量
	err error
)

func init() {
	options := Constant2.SqlOptions{
		Ip:     "127.0.0.1",
		Port:   "3306",
		User:   "root",
		Pwd:    "12345",
		DBName: "finbook_demo",
		DBType: Constant2.MySql,
	}

	Db, err = Prodviders.CreateDB(&options)
	if err != nil {
		panic(err)
	}

}

func Exec(str string, args ...any) (int64, error) {

	if len(str) <= 0 {
		return 0, errors.New("参数不能为空" + err.Error())
	}

	//预编译sql
	stmt, err := Db.Prepare(str)
	//defer Db.Close()
	if err != nil {
		return 0, err
	}

	defer stmt.Close() // 确保在函数返回前关闭语句
	// 执行
	ret, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	rows, err := ret.RowsAffected()
	return rows, err
}

func Query(str string, args ...any) (*sql.Rows, error) {

	if len(str) <= 0 {
		return nil, errors.New("参数不能为空")
	}

	//预编译sql
	stmt, err := Db.Prepare(str)
	//defer Db.Close()
	if err != nil {
		return nil, errors.New("预编译sql失败" + err.Error())
	}

	defer stmt.Close() // 确保在函数返回前关闭语句
	// 执行
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	return rows, err

}

func QueryRow(str string, args ...any) (*sql.Row, error) {

	if len(str) <= 0 {
		return nil, errors.New("参数不能为空")
	}

	//预编译sql
	stmt, err := Db.Prepare(str)
	//defer Db.Close()
	if err != nil {
		return nil, errors.New("预编译sql失败" + err.Error())
	}

	defer stmt.Close() // 确保在函数返回前关闭语句
	// 执行
	row := stmt.QueryRow(args...) // 使用 ... 来展开
	if err != nil {
		return nil, err
	}

	return row, err

}
