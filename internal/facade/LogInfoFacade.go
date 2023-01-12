package facade

import (
	"regexp"
	"strings"
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
}

func NewLogInfoFacade(
	logInfoService logInfoService,
	readLogService readLogService,
	regexService regexService,
	scanLogService scanLogService,
	validationLogService validationLogService,
) *logInfoFacade {
	return &logInfoFacade{
		logInfoService:       logInfoService,
		readLogService:       readLogService,
		regexService:         regexService,
		scanLogService:       scanLogService,
		validationLogService: validationLogService,
	}
}

func (l *logInfoFacade) ProcessingLogs() error {
	re := regexp.MustCompile(pattern)
	regex = re

	file, err := l.readLogService.ReadFile("logs")
	if err != nil {
		return err
	}
	defer file.Close()

	chanRow := l.scanLogService.ScanFile(file)

	for logRow := range chanRow {

		values, err := l.regexService.SeparateByGroups(logRow)
		if err != nil {
			return err
		}

		if !l.validationLogService.ValidateRow(*values) {
			return err
		}

		strings.Split()

	}

	//validar grupos

	//inteirar no array

	//persistir na base

}
