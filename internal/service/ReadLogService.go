package service

import (
	"fmt"
	"os"
)

type readLogService struct {
}

func NewReadLogService() *readLogService {
	return &readLogService{}
}

func (r *readLogService) ReadFile(path string) (*os.File, error) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	return file, nil
}
