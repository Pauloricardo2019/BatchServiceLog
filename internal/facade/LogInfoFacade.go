package facade

import "regexp"

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

	defer file.Close()

	//processar linha por linha

	//range no channel

	//processar regex

	//validar grupos

	//inteirar no array

	//persistir na base

}
