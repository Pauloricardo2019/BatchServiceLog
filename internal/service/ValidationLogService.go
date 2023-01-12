package service

type validationLogService struct {
}

func NewValidationLogService() *validationLogService {
	return &validationLogService{}
}

func (v *validationLogService) ValidateRow(log string) bool {

}
