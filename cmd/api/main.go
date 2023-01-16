package main

import (
	"batch-service/internal/facade"
	"batch-service/internal/model"
	"batch-service/internal/provider"
	"batch-service/internal/repository"
	"batch-service/internal/service"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := service.NewGetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{
		SkipDefaultTransaction: true,
		FullSaveAssociations:   true,
	})

	db.AutoMigrate(model.Log{})

	if err != nil {
		panic(err)
	}

	logRepository := repository.NewLogInfo(db)

	logProvider := provider.NewLogWriteProvider()

	readLogService := service.NewReadLogService()
	regexService := service.NewRegexService()
	scanLogService := service.NewScanLogService()
	validationWebLogService := service.NewValidationLogService(logProvider)
	webLogInfoService := service.NewLogInfoService(logRepository)

	webLogFacade := facade.NewLogInfoFacade(webLogInfoService, readLogService, regexService, scanLogService, validationWebLogService, logProvider)

	err = webLogFacade.ProcessingLogs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Finish !")

}
