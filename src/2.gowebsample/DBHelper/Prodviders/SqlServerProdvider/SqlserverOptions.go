package SqlServerProdvider

import (
	"2.gowebsample/Common/Constant"
	"fmt"
)

func Init(options *Constant.SqlOptions) string {
	//dsn := "server=localhost;user id=username;password=password;database=dbname"
	//db, err := sql.Open("sqlserver", dsn)

	dsn := fmt.Sprintf("server=%v:%v;user id=%v;password=%v;database=%v",
		options.Ip,
		options.Port,
		options.User,
		options.Pwd,
		options.DBName)
	return dsn
}
