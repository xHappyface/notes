package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xHappyface/notes/log"
)

type NotesDB struct {
	source *sql.DB
	ctx    context.Context
}

var (
	errDBNotExists = errors.New(`Please read "README.md" for instructions on setting up the database.`)
)

func (db *NotesDB) Close() {
	db.source.Close()
}

func LoadDB(logger *log.NotesLogger, ctx context.Context) (*NotesDB, error) {
	logger.Log(log.LOG_LV_INFO, "Loading database...")
	logger.Log(log.LOG_LV_INFO, "Connecting to database")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/notes", os.Getenv("DB_USER"), os.Getenv("DB_PASS")))
	if err != nil {
		return nil, err
	}
	logger.Log(log.LOG_LV_INFO, "Connected to database")
	logger.Log(log.LOG_LV_INFO, "Pinging database")
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Join(err, errDBNotExists)
	}
	logger.Log(log.LOG_LV_INFO, "Database pinged")
	logger.Log(log.LOG_LV_INFO, "Environment loaded successfully!")
	return &NotesDB{
		source: db,
		ctx:    ctx,
	}, nil
}

func (db *NotesDB) InitDB(logger *log.NotesLogger) error {
	logger.Log(log.LOG_LV_INFO, "Initializing database...")
	logger.Log(log.LOG_LV_INFO, "Preparing statement")
	stmt, err := db.source.PrepareContext(db.ctx, `CREATE TABLE IF NOT EXISTS notes(id VARCHAR(36) PRIMARY KEY NOT NULL, title VARCHAR(180) NOT NULL DEFAULT "", body TEXT);`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	logger.Log(log.LOG_LV_INFO, "Statement prepared")
	logger.Log(log.LOG_LV_INFO, "Executing statement")
	_, err = stmt.ExecContext(db.ctx)
	if err != nil {
		return err
	}
	logger.Log(log.LOG_LV_INFO, "Statement executed")
	logger.Log(log.LOG_LV_INFO, "Database initialized successfully!")
	return nil
}
