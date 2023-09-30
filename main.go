package main

import (
	"github.com/joho/godotenv"
	"github.com/xHappyface/notes/log"
)

func main() {
	logger := log.NewLogger()
	if err := godotenv.Load(); err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error())
	}
}
