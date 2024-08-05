package store

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"main/backend/config"
	"strings"

	"github.com/go-redis/redis/v8"
)

var (
	db  *sql.DB
	rdb *redis.Client
)

type RedisNamespace string

const (
	redisNamespaceSession     RedisNamespace = "session"
	redisNamespaceAccessToken RedisNamespace = "access_token"
)

func redisKey(args ...string) string {
	return strings.Join(args, ":")
}

func Init() error {
	// connect to postgres
	var err error
	db, err = sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		slog.Error("Failed to connect to Postgres")
		return err
	}
	slog.Info("successfully initialize Postgres")

	// connect to Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		slog.Error("Failed to connect Redis")
		return err
	}
	slog.Info("Successfully initialize Redis")
	return nil
}

func Close() (error, error) {
	return func() error {
			err := db.Close()
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to close db: %s", err.Error()))
				return err
			}
			slog.Info("Postgres connection has been closed")
			return nil
		}(), func() error {
			err := rdb.Close()
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to close rdb: %s", err.Error()))
				return err
			}
			slog.Info("Redis connection has been closed")
			return nil
		}()
}
