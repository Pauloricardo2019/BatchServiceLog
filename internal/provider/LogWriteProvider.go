package provider

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var logMux sync.Mutex

const logPath string = "./log.txt"

type logWriteProvider struct {
}

func NewLogWriteProvider() *logWriteProvider {
	return &logWriteProvider{}
}

func (l *logWriteProvider) LogInfo(text string) {
	l.writeLog("INFO", text)
}

func (l *logWriteProvider) LogError(text string) {
	l.writeLog("INFO", text)
}

func (l *logWriteProvider) writeLog(text string, logType string) error {
	logMux.Lock()
	defer logMux.Unlock()

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	now := time.Now().Format(time.RFC3339)
	logLine := fmt.Sprintf("[%s] [%s] - %s\n", now, logType, text)
	_, err = f.WriteString(logLine)

	if err != nil {
		log.Println(err)
	}
	return nil
}
