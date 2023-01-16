package service

import (
	"batch-service/internal/model"
	"context"
)

type logInfoService struct {
	logInfoRepository logInfoRepository
}

func NewLogInfoService(logInfoRepository logInfoRepository) *logInfoService {
	return &logInfoService{
		logInfoRepository: logInfoRepository,
	}
}

func (l *logInfoService) InsertLogInfo(ctx context.Context, log *model.Log) error {
	return l.logInfoRepository.InsertRecords(ctx, log)
}
