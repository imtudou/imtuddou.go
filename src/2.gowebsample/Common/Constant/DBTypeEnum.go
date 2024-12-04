package Constant

type DBType int

const (
	MySql DBType = iota
	PostgreSQL
	Sqlite
	SqlServer
)

func (dbType DBType) DriverName() string {
	switch dbType {
	case MySql:
		return "mysql"
	case PostgreSQL:
		return "postgres" // 注意 PostgreSQL 的驱动名称通常是 "postgres" 而不是 "postgresql"
	case Sqlite:
		return "sqlite3" // SQLite 的驱动名称通常是 "sqlite3"
	case SqlServer:
		return "sqlserver" // SQL Server 的驱动名称可能因使用的驱动而异
	default:
		return ""
	}
}
