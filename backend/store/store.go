package store

import (
	"database/sql"
	"fmt"
	"log/slog"
	"main/backend/config"
)

var (
	db *sql.DB
)

func Init() error {
	//connect to postgres
	var err error
	db, err = sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		slog.Error("Failed to connect to postgre database")
		return err
	}
	slog.Info("successfully initialize postgres")
	return nil
}

func Close() error {
	err := db.Close()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to close db: %s", err.Error()))
		return err
	}
	slog.Info("db has been closed")
	return nil
}
