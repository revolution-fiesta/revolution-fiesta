package store

import (
	"fmt"
	"log/slog"
	"main/backend/config"
	"os"

	_ "github.com/lib/pq"
)

// 迁移.
func Migrate() error {
	// 输出当前版本号到日志
	slog.Info(config.Version)

	// 检查info表是否存在
	var info_exists bool
	err := db.QueryRow(`SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'public'
		AND table_name = 'info'
	);`).Scan(&info_exists)
	if err != nil {
		slog.Error("Failed to check if info table exists")
	}

	if !info_exists {
		schemaSql, err := os.ReadFile(config.SchemaFilePath)
		if err != nil {
			slog.Error("Failed to read schema file")
			return err
		}

		// 执行 schema.sql 文件中的SQL语句
		_, err = db.Exec(string(schemaSql))
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to execute schema file: %s", err.Error()))
			return err
		}
		slog.Info("Database schema has updated successfully")
		return nil
	}

	slog.Info("Skip migration")
	return nil
}
