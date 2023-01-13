package repository

import (
	"batch-service/internal/model"
	"context"
	"gorm.io/gorm"
)

type logInfoRepository struct {
	*BaseRepository
}

func NewLogInfo(db *gorm.DB) *logInfoRepository {
	baseRepo := NewBaseRepository(db)
	return &logInfoRepository{
		baseRepo,
	}
}

func (l *logInfoRepository) InsertRecords(ctx context.Context, logs []model.Log) error {
	db, err := l.GetConnection(ctx)
	if err != nil {
		return err
	}

}
