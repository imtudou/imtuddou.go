package SqliteProvider

import (
	"2.gowebsample/Common/Constant"
)

func Init(options *Constant.SqlOptions) string {

	//dsn := "./dbname.db" // SQLite数据库文件路径
	//db, err := sql.Open("sqlite3", dsn)

	// 直接传路径到 DBName 即可
	return options.DBName
}
