package store

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"main/backend/config"

	"github.com/go-redis/redis/v8"
)

var (
	db  *sql.DB
	rdb *redis.Client
)

func Init() error {
	// connect to postgres
	var err error
	db, err = sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		slog.Error("Failed to connect to Postgre database")
		return err
	}
	slog.Info("successfully initialize Postgres")

	// connect to Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		slog.Error("Failed to connect Redis")
		return err
	}
	slog.Info("successfully initialize Redis")
	return nil
}

func Close() (error, error) {
	return func() error {
			err := db.Close()
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to close db: %s", err.Error()))
				return err
			}
			slog.Info("db has been closed")
			return nil
		}(), func() error {
			err := rdb.Close()
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to close rdb: %s", err.Error()))
				return err
			}
			slog.Info("rdb has been closed")
			return nil
		}()
}
