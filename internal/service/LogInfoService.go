package service

import (
	"batch-service/internal/model"
	"context"
)

type logInfoRepository interface {
	InsertRecords(ctx context.Context, logs []model.Log) error
}

type logInfoService struct {
	logInfoRepository logInfoRepository
}

func NewLogInfoService(logInfoRepository logInfoRepository) *logInfoService {
	return &logInfoService{
		logInfoRepository: logInfoRepository,
	}
}

func (l *logInfoService) InsertLogInfo(ctx context.Context, logs []model.Log) error {
	return l.logInfoRepository.InsertRecords(ctx, logs)
}
