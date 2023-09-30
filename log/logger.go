package log

import (
	"log"
	"os"
)

const (
	LOG_LV_INFO uint8 = iota + 1
	LOG_LV_WRN
	LOG_LV_ERR
	LOG_LV_FTL_ERR

	PREFIX_LV_INFO    = "INFO"
	PREFIX_LV_WRN     = "WRN"
	PREFIX_LV_ERR     = "ERR"
	PREFIX_LV_FTL_ERR = "FATAL_ERR"
)

type NotesLogger struct {
	output *log.Logger
}

func NewLogger() *NotesLogger {
	return &NotesLogger{
		output: log.New(os.Stderr, "NOTES:", log.LstdFlags|log.Lmsgprefix),
	}
}

func (logger *NotesLogger) Log(lv uint8, msg string) {
	switch lv {
	case LOG_LV_INFO:
		logger.output.Printf("%s:%s", PREFIX_LV_INFO, msg)
	case LOG_LV_WRN:
		logger.output.Printf("%s:%s", PREFIX_LV_WRN, msg)
	case LOG_LV_ERR:
		logger.output.Printf("%s:%s", PREFIX_LV_ERR, msg)
	case LOG_LV_FTL_ERR:
		logger.output.Fatalf("%s:%s", PREFIX_LV_FTL_ERR, msg)
	}
}
