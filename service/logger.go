package service

import (
	"time"

	"devops.inspur.com/ITE__InTech-blockchain/logmodule/models"
	"devops.inspur.com/ITE__InTech-blockchain/logmodule/storage"
)

type LoggerService struct{}

func NewLoggerService() *LoggerService {
	return &LoggerService{}
}

// 记录日志
func (s *LoggerService) RecordLog(description, requestType, requestPath, className, methodName, initiator, status string) error {
	return storage.AddLogEntry(description, requestType, requestPath, className, methodName, initiator, status)
}

// 获取所有日志记录
func (s *LoggerService) GetAllLogs() ([]models.LogEntry, error) {
	return storage.GetAllLogs()
}

// 根据日志说明关键字检索日志
func (s *LoggerService) SearchLogsByDescription(keyword string) ([]models.LogEntry, error) {
	return storage.SearchLogsByDescription(keyword)
}

// 根据指定的时间范围检索日志
func (s *LoggerService) GetLogsByTimeRange(startTime, endTime time.Time) ([]models.LogEntry, error) {
	return storage.GetLogsByTimeRange(startTime, endTime)
}

// 联合检索
func (s *LoggerService) RetrieveLogs(keyword string, startTime, endTime time.Time) ([]models.LogEntry, error) {
	return storage.RetrieveLogs(keyword, &startTime, &endTime)
}
