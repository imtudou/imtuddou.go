package Prodviders

import (
	Constant2 "2.gowebsample/Common/Constant"
	"2.gowebsample/DBHelper/Prodviders/MySqlProdvider"
	"2.gowebsample/DBHelper/Prodviders/PostgreSQLProdvider"
	"2.gowebsample/DBHelper/Prodviders/SqlServerProdvider"
	"2.gowebsample/DBHelper/Prodviders/SqliteProvider"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDB(options *Constant2.SqlOptions) (*sql.DB, error) {
	if options == nil {
		return nil, errors.New("【CreateDB】.SqlOptions 参数不能为空！")
	}

	var dsn string
	switch options.DBType {
	case Constant2.MySql:
		dsn = MySqlProdvider.Init(options) // 需要安装实现包：go get -u github.com/go-sql-driver/mysql
		break
	case Constant2.PostgreSQL:
		dsn = PostgreSQLProdvider.Init(options) // 需要安装实现包
		break
	case Constant2.Sqlite:
		dsn = SqliteProvider.Init(options) // 需要安装实现包
		break
	case Constant2.SqlServer:
		dsn = SqlServerProdvider.Init(options) // 需要安装实现包
		break
	default:
		dsn = MySqlProdvider.Init(options) // 需要安装实现包
		break
	}

	db, err := sql.Open(options.DBType.DriverName(), dsn)
	if err != nil {
		return nil, err // 返回错误而不是 panic
	}

	// 可选：检查数据库连接是否有效
	if err := db.Ping(); err != nil {
		db.Close() // 如果 Ping 失败，关闭数据库连接
		return nil, err
	}

	return db, err
}
