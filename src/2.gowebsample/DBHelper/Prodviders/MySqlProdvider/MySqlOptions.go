package MySqlProdvider

import (
	"2.gowebsample/Common/Constant"
	"fmt"
)

func Init(options *Constant.SqlOptions) string {

	//dsn := "username:password@tcp(127.0.0.1:3306)/dbname"
	//db, err := sql.Open("mysql", dsn)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		options.User,
		options.Pwd,
		options.Ip,
		options.Port,
		options.DBName)
	return dsn
}
