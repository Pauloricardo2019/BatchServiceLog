package service

import (
	"batch-service/internal/model"
)

var (
	verbs = []string{
		"POST",
		"PUT",
		"GET",
		"DELETE",
		"HEAD",
	}
)

type validationLogService struct {
	logProvider logProvider
}

func NewValidationLogService(logProvider logProvider) *validationLogService {
	return &validationLogService{
		logProvider: logProvider,
	}
}

func (v *validationLogService) ValidateRow(log *model.Log) bool {

	if len(log.IP) > 15 {
		v.logProvider.LogInfo("This ip is a IPV6!!!")
	}

	if !v.validHttpVerbs(log.Verb) {
		v.logProvider.LogError("Invalid http verb")
		return false
	}

	return true
}

func (v *validationLogService) validHttpVerbs(logVerb string) bool {
	for _, verb := range verbs {
		if logVerb == verb {
			return true
		}

	}
	return false
}
