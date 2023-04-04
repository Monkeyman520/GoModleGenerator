package error

var (
	InvalidDatabaseHost     = NewPreDefineZhCNInternalError("invalid database host", "域名/ip无效")
	InvalidDatabasePort     = NewPreDefineZhCNInternalError("invalid database port", "数据库端口无效")
	InvalidDatabaseUsername = NewPreDefineZhCNInternalError("invalid database username", "数据库用户名无效")
	InvalidDatabasePassword = NewPreDefineZhCNInternalError("invalid database password", "数据库密码无效")
	InvalidDatabaseSchema   = NewPreDefineZhCNInternalError("invalid database schema", "数据库Schema无效")

	DbConnectionError = NewPreDefineZhCNInternalError("db connection error", "数据库连接失败")
	ViperReadError    = NewPreDefineZhCNInternalError("viper read error", "viper读取配置失败")

	TablesSearchError = NewPreDefineZhCNInternalError("search table error", "table筛选失败")

	CheckPathError = NewPreDefineZhCNInternalError("check path error", "路径检查失败")
	MakePathError  = NewPreDefineZhCNInternalError("create path error", "创建路径失败")

	UnknowPathError = NewPreDefineZhCNInternalError("check path end with unknown error", "路径检查未知错误")
)
