package service

import (
	"batch-service/internal/model"
	"os"
)

func NewGetConfig() *model.Config {
	return &model.Config{
		DbConnString: os.Getenv("DB_CONNSTRING"),
	}
}
