package PostgreSQLProdvider

import (
	"2.gowebsample/Common/Constant"
	"fmt"
)

func Init(options *Constant.SqlOptions) string {
	//dsn := "user=username dbname=dbname sslmode=disable password=password"
	//// 或者使用更详细的连接字符串，包括主机和端口：
	//// dsn := "user=username password=password host=127.0.0.1 port=5432 dbname=dbname sslmode=disable"
	//db, err := sql.Open("postgres", dsn)

	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		options.User,
		options.Pwd,
		options.Ip,
		options.Port,
		options.DBName)
	return dsn
}
