package migration

import (
	"log/slog"
	"main/backend/config"
)

// 迁移.
func Migrate() {
	// 你好，把这个函数实现一下.
	// 把这个 config.Version 变量存到 info 表里.
	slog.Info(config.Version)
	// 如果检查到存在这张表并且有这个变量，就运行 schema.sql 里的 SQL 语句自动构建数据库.
	// 如果没有说明数据库没构建过，跳过这个函数就好.
}
