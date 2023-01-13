package service

var (
	verbs = []string{
		"POST",
		"PUT",
		"GET",
		"DELETE",
	}
)

type validationLogService struct {
}

func NewValidationLogService() *validationLogService {
	return &validationLogService{}
}

func (v *validationLogService) ValidateRow(log string) bool {
	if !v.validHttpVerbs() {
		
	}

}

func (v *validationLogService) validHttpVerbs(logVerb string) bool {
	for _, verb := range verbs {
		if logVerb == verb {
			return true
		}

	}
	return false
}