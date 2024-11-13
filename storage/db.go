package storage

import (
	"fmt"

	"devops.inspur.com/ITE__InTech-blockchain/logmodule/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error

	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	// MySQL 自动迁移，创建表并确保字符集为 utf8mb4
	// err = DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&models.LogEntry{})
	// if err != nil {
	// 	return fmt.Errorf("自动迁移失败: %v", err)
	// }

	// pg 自动迁移
	if err := DB.AutoMigrate(&models.LogEntry{}).Error; err != nil {
		return fmt.Errorf("自动迁移失败: %v", err)
	}

	return nil
}
