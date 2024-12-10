package storage

import (
	"encoding/hex"
	"time"

	"logmodule/utils"

	"errors"

	"devops.inspur.com/ITE__InTech-blockchain/logmodule/models"
	"github.com/EthanWalker10/logmodule/utils"
)

func GetAllLogs() ([]models.LogEntry, error) {
	var logs []models.LogEntry
	err := DB.Find(&logs).Error
	return logs, err
}

func AddLog(entry *models.LogEntry) error {
	return DB.Create(entry).Error
}

func AddLogEntry(description, requestType, requestPath, className, methodName, initiator, status string) error {
	data := []byte(description + requestType + requestPath + className + methodName + initiator + status)
	hash, err := utils.Sha256(data)
	if err != nil {
		return "", err
	}
	hashHex := hex.EncodeToString(hash)
	entry := models.LogEntry{
		Description: description,
		RequestType: requestType,
		RequestPath: requestPath,
		ClassName:   className,
		MethodName:  methodName,
		RequestTime: time.Now(),
		Initiator:   initiator,
		Status:      status,
		hash:        hashHex,
	}
	return AddLog(&entry)
}

func SearchLogsByDescription(keyword string) ([]models.LogEntry, error) {
	var logs []models.LogEntry
	err := DB.Where("description LIKE ?", "%"+keyword+"%").Find(&logs).Error
	return logs, err
}

func GetLogsByTimeRange(startTime, endTime time.Time) ([]models.LogEntry, error) {
	var logs []models.LogEntry
	err := DB.Where("request_time BETWEEN ? AND ?", startTime, endTime).Find(&logs).Error
	return logs, err
}

// GetLogs 根据参数获取日志记录
// - 如果 keyword 不为空，则进行关键字查询
// - 如果 startTime 和 endTime 不为空，则在时间区间内查询
// - 如果均为空，则返回所有日志记录
// GetLogs 综合查询日志，支持按关键字、时间区间筛选
func RetrieveLogs(keyword string, startTime, endTime *time.Time) ([]models.LogEntry, error) {
	var logs []models.LogEntry
	query := DB.Model(&models.LogEntry{})

	// 根据 keyword 增加模糊搜索条件
	if keyword != "" {
		query = query.Where("description LIKE ?", "%"+keyword+"%")
	}

	// 根据时间区间增加查询条件
	if startTime != nil && endTime != nil {
		query = query.Where("request_time BETWEEN ? AND ?", *startTime, *endTime)
	}

	// 执行查询
	err := query.Find(&logs).Error
	return logs, err
}

func GetLogAndValidateHash(id string) (*models.LogEntry, error) {
	var entry models.LogEntry
	err := DB.First(&entry, id).Error
	if err != nil {
		return nil, err
	}

	data := []byte(entry.Description + entry.RequestType + entry.RequestPath + entry.ClassName + entry.MethodName + entry.Initiator + entry.Status)
	hash, err := utils.Sha256(data)
	if err != nil {
		return nil, err
	}
	hashHex := hex.EncodeToString(hash)
	if entry.Hash != hashHex {
		return nil, errors.New("stringhash mismatch for log entry with id %s")
	}

	return &entry, nil
}
