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

	// 另外数据库连接的那些变量可以和 config.Version 放在一起, 每次启动服务器前去那里填写好启动程序就好了
	// 位置在 -> backend\config\config.go
}
