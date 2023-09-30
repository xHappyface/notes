package main

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/xHappyface/notes/db"
	"github.com/xHappyface/notes/log"
)

func main() {
	logger := log.NewLogger()
	logger.Log(log.LOG_LV_INFO, "Loading environment")
	if err := godotenv.Load(); err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error())
	}
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()
	db, err := db.LoadDB(logger, ctx)
	if err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error())
	}
	defer db.Close()
	if err = db.InitDB(logger); err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error())
	}
}
