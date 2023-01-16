package service

import (
	"batch-service/internal/model"
	"context"
)

type logProvider interface {
	LogInfo(text string)
	LogError(text string)
}

type logInfoRepository interface {
	InsertRecords(ctx context.Context, log *model.Log) error
}
