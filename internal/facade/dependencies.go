package facade

import (
	"batch-service/internal/model"
	"context"
	"os"
	"regexp"
)

type logInfoService interface {
	InsertLogInfo(ctx context.Context, logs *model.Log) error
}

type readLogService interface {
	ReadFile(path string) (*os.File, error)
}

type regexService interface {
	SeparateByGroups(regex *regexp.Regexp, logRow string) (*model.Log, error)
}

type scanLogService interface {
	ScanFile(file *os.File) chan string
}

type validationLogService interface {
	ValidateRow(log *model.Log) bool
}

type logProvider interface {
	LogInfo(text string)
	LogError(text string)
}
