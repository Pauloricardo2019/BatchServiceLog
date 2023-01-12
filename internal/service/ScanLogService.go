package service

import (
	"bufio"
	"io"
	"os"
)

type scanLogService struct {
}

func NewScanLogService() *scanLogService {
	return &scanLogService{}
}

func (r *scanLogService) ScanFile(file *os.File) chan string {

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	result := make(chan string)

	go func() {
		for scanner.Scan() {

			logRow := scanner.Text()
			if scanner.Err() == io.EOF {
				return
			}

			result <- logRow

		}
		if scanner.Err() == io.EOF {
			return
		}
	}()
	return result
}
