package config

import (
	"os"
)

// DBConfig 数据库配置信息
type DBConfig struct {
	DSN string // 数据库连接
}

// LoadConfig 从环境变量或默认值加载配置
func LoadConfig() *DBConfig {
	dsn := os.Getenv("LOGMODULE_DB_DSN")
	if dsn == "" {
		// dsn = "user:password@tcp(localhost:3306)/logdb" // MySQL 本地测试容器
		dsn = "chainmaker:Q@su!6vGm9@tcp(192.168.139.67:5866)" // pg服务器
	}
	return &DBConfig{DSN: dsn}
}
