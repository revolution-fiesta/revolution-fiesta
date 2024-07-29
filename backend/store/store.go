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
	Db  *sql.DB
	Rdb *redis.Client
	Ctx context.Context
)

func Init() error {
	// connect to postgres
	var err error
	Db, err = sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		slog.Error("Failed to connect to Postgre database")
		return err
	}
	slog.Info("successfully initialize Postgres")

	// connect to Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	Ctx = context.Background()
	_, err = Rdb.Ping(Ctx).Result()
	if err != nil {
		slog.Error("Failed to connect Redis")
		return err
	}
	slog.Info("successfully initialize Redis")
	return nil
}

func Close() (error, error) {
	return func() error {
			err := Db.Close()
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to close db: %s", err.Error()))
				return err
			}
			slog.Info("db has been closed")
			return nil
		}(), func() error {
			err := Rdb.Close()
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to close rdb: %s", err.Error()))
				return err
			}
			slog.Info("rdb has been closed")
			return nil
		}()
}
