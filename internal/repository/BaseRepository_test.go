package repository

import (
	"batch-service/internal/service"
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestBeginTransaction(t *testing.T) {
	cfg := service.NewGetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString))
	if err != nil {
		panic(err.Error())
	}

	baseRepo := NewBaseRepository(db)

	ctx := context.Background()
	_, err = baseRepo.BeginTransaction(ctx)
	assert.NoError(t, err)
}

func TestCommitTransactionOK(t *testing.T) {
	cfg := service.NewGetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString))
	if err != nil {
		panic(err.Error())
	}

	baseRepo := NewBaseRepository(db)

	ctx := context.Background()
	newCtx, err := baseRepo.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = baseRepo.CommitTransaction(newCtx)
	assert.NoError(t, err)
}

func TestCommitTransactionNoBeginTransaction(t *testing.T) {
	cfg := service.NewGetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString))
	if err != nil {
		panic(err.Error())
	}

	baseRepo := NewBaseRepository(db)

	ctx := context.Background()
	err = baseRepo.CommitTransaction(ctx)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "no transaction found on current context")
}
