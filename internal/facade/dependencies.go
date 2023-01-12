package facade

import (
	"batch-service/internal/model"
	"context"
	"os"
)

type logInfoService interface {
	InsertLogInfo(ctx context.Context, logs []model.Log) error
}

type readLogService interface {
	ReadFile(path string) (*os.File, error)
}

type regexService interface {
	SeparateByGroups(logRow string) (*string, error)
}

type scanLogService interface {
	ScanFile(file *os.File) chan string
}

type validationLogService interface {
	ValidateRow(log string) bool
}
