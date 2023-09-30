package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/xHappyface/notes/log"
)

func main() {
	logger := log.NewLogger()
	logger.Log(log.LOG_LV_INFO, "Loading environment")
	if err := godotenv.Load(); err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error())
	}
	logger.Log(log.LOG_LV_INFO, "Connecting to database")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/notes", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))
	if err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error())
	}
	logger.Log(log.LOG_LV_INFO, "Pinging database")
	if err = db.Ping(); err != nil {
		logger.Log(log.LOG_LV_FTL_ERR, err.Error()+". Possible cause of error is database not existing. Please read \"README.md\" for instructions on setting up the database.")
	}
}
