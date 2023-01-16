package repository

import (
	"batch-service/internal/model"
	"context"
	"fmt"
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

func (l *logInfoRepository) InsertRecords(ctx context.Context, log *model.Log) error {
	db, err := l.GetConnection(ctx)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO logs (ip, date, verb) VALUES ('%s', '%s', '%s'); ", log.IP, log.Date, log.Verb)

	return db.Exec(query).Error
}
