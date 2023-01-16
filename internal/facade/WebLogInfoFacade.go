package facade

import (
	"context"
	"fmt"
	"regexp"
)

var regex *regexp.Regexp

const (
	pattern = "((?P<IPV4>(?:[0-9]{1,3}\\.){3}[0-9]{1,3})|(?P<IPV6>(?:[0-9a-z]+\\:){7}[0-9a-z]+))\\s\\-\\s\\-\\s(?P<date>\\[[0-9]+\\/[a-zA-Z]+\\/[0-9]+(?:\\:[0-9]+){3}\\s\\+[0-9]+\\])\\s\\\"(?P<Verb>[A-Z]+)*"
)

type logInfoFacade struct {
	logInfoService       logInfoService
	readLogService       readLogService
	regexService         regexService
	scanLogService       scanLogService
	validationLogService validationLogService
	logProvider          logProvider
}

func NewLogInfoFacade(
	logInfoService logInfoService,
	readLogService readLogService,
	regexService regexService,
	scanLogService scanLogService,
	validationLogService validationLogService,
	logProvider logProvider,
) *logInfoFacade {
	return &logInfoFacade{
		logInfoService:       logInfoService,
		readLogService:       readLogService,
		regexService:         regexService,
		scanLogService:       scanLogService,
		validationLogService: validationLogService,
		logProvider:          logProvider,
	}
}

func (l *logInfoFacade) ProcessingLogs() error {
	ctx := context.Background()
	l.logProvider.LogInfo("Start processing")
	re := regexp.MustCompile(pattern)
	regex = re

	file, err := l.readLogService.ReadFile("logs")
	if err != nil {
		l.logProvider.LogError("read file error: " + err.Error())
		return err
	}
	defer file.Close()

	chanRow := l.scanLogService.ScanFile(file)

	for logRow := range chanRow {
		l.logProvider.LogInfo("Start range on channel")
		log, err := l.regexService.SeparateByGroups(regex, logRow)
		if err != nil {
			l.logProvider.LogError("regex error: " + err.Error())
			continue
		}

		l.logProvider.LogInfo("passed by separate by groups")

		l.logProvider.LogInfo(fmt.Sprintf("Log ip %s, date %s, verb %s.", log.IP, log.Date, log.Verb))

		if !l.validationLogService.ValidateRow(log) {
			l.logProvider.LogError("validate error: " + err.Error())
			continue
		}

		l.logProvider.LogInfo("passed by validation")

		err = l.logInfoService.InsertLogInfo(ctx, log)
		if err != nil {
			l.logProvider.LogError("insert database error: " + err.Error())
			continue
		}

		l.logProvider.LogInfo("finish processing")
	}
	return nil
}
