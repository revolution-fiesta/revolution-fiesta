package migration

import (
	"database/sql"
	"log/slog"
	"main/backend/config"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

// 迁移.
func Migrate() {

	// 读取配置信息
	connStr := config.DatabaseUrl
	version := config.Version // 获取当前版本号

	// 输出当前版本号到日志
	slog.Info(config.Version)

	// 连接到数据库
	db, err := sql.Open("postgres", connStr)
	// 如果数据库连接失败，记录错误并返回
	if err != nil {
		slog.Error("Failed to connect to postgre database", err)
		return
	}
	defer db.Close() // 确保函数结束时关闭数据库连接

	// 检查info表是否存在
	var info_exists bool
	err = db.QueryRow(`SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'public'
		AND table_name = 'info'
	);`).Scan(&info_exists)

	// 如果检查info表是否存在失败，返回失败并记录到日志
	if err != nil {
		slog.Error("Failed to check if info table exists", err)
	}

	// 如果info表存在的话
	if info_exists {
		// 检查一下版本号有没有在info表上
		var version_exists string // version_exists 含义是看看版本号是否有在info表上
		err = db.QueryRow(`SELECT version FROM info WHERE version = $1`, version).Scan(&version_exists)

		// 如果查询失败并且错误不是"没有记录"的话，返回并记录错误到日志上
		if err != nil && err != sql.ErrNoRows {
			slog.Error("Fail to query version", err)
			return
		}

		// 如果info表上不存在该版本的记录，则进行插入
		if version_exists == "" {
			_, err = db.Exec(`INSERT INTO info (version) VALUES ($1)`, version)

			// 如果插入失败，返回并把错误记录到日志上
			if err != nil {
				slog.Error("Failed to insert version into info table", err)
				return
			}
		}

		// 执行 schema.sql 文件中的SQL语句
		schemaPath := filepath.Join("main", "backend", "schema.sql")
		schemaSQL, err := os.ReadFile(schemaPath)
		// 如果文件读取失败，返回并把错误记录到日志上
		if err != nil {
			slog.Error("Failed to read schema.sql", err)
			return
		}
		_, err = db.Exec(string(schemaSQL)) // 执行SQl语句
		if err != nil {
			slog.Error("Failed to execute schema.sql", err) //失败返回记录日志
			return
		}
		slog.Info("Database schema updated")
	}

}
