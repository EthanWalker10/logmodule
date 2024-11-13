package models

import "time"

// MySQL 数据库表结构
// type LogEntry struct {
// 	ID          int       `json:"id" gorm:"primaryKey"`                  // 日志的唯一标识，主键
// 	Description string    `json:"description" gorm:"type:text"`          // 日志说明
// 	RequestType string    `json:"request_type" gorm:"type:varchar(50)"`  // 请求类型
// 	RequestPath string    `json:"request_path" gorm:"type:varchar(255)"` // 请求路径
// 	ClassName   string    `json:"class_name" gorm:"type:varchar(100)"`   // 请求的类名
// 	MethodName  string    `json:"method_name" gorm:"type:varchar(100)"`  // 请求的方法名
// 	RequestTime time.Time `json:"request_time" gorm:"type:datetime"`     // 请求发生的时间
// 	Initiator   string    `json:"initiator" gorm:"type:varchar(100)"`    // 请求发起人
// 	Status      string    `json:"status" gorm:"type:varchar(20)"`        // 操作状态
// }

// pg 数据库表结构
type LogEntry struct {
	ID          int       `json:"id" gorm:"primaryKey"`                  // 日志的唯一标识，主键
	Description string    `json:"description" gorm:"type:text"`          // 日志说明
	RequestType string    `json:"request_type" gorm:"type:varchar(50)"`  // 请求类型
	RequestPath string    `json:"request_path" gorm:"type:varchar(255)"` // 请求路径
	ClassName   string    `json:"class_name" gorm:"type:varchar(100)"`   // 请求的类名
	MethodName  string    `json:"method_name" gorm:"type:varchar(100)"`  // 请求的方法名
	RequestTime time.Time `json:"request_time" gorm:"type:timestamp"`    // 请求发生的时间
	Initiator   string    `json:"initiator" gorm:"type:varchar(100)"`    // 请求发起人
	Status      string    `json:"status" gorm:"type:varchar(20)"`        // 操作状态
}
