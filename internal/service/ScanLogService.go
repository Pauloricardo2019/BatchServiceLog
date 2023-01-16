package service

import (
	"bufio"
	"io"
	"os"
	"time"
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
		time.Sleep(time.Second * 1)
		for scanner.Scan() {

			logRow := scanner.Text()
			if logRow == "" || scanner.Err() == io.EOF {
				close(result)
				return
			}

			result <- logRow
		}
	}()
	return result
}
