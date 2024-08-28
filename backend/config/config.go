package config

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"
)

const (
	Version               string        = "0.0.1"
	Port                  string        = "8080"
	DatabaseUsr           string        = "postgres"
	DatabasePasswd        string        = "postgres"
	DatabaseHost          string        = "114.132.70.108"
	DatabasePort          string        = "5432"
	DatabaseName          string        = "postgres"
	SchemaFilePath        string        = "../schema.sql"
	RedisAddr             string        = "114.132.70.108:6379"
	RedisPassword         string        = "redis"
	RedisDB                             = 0
	ProjectName           string        = "revolution-fiesta"
	AccessTokenExpiration time.Duration = time.Hour / 4
)

var (
	DatabaseUrl string
	LocalAddr   string
	PrivateKey  *rsa.PrivateKey
)

func init() {
	DatabaseUrl = getPgConnUrl(DatabaseUsr, DatabasePasswd, DatabaseHost, DatabasePort, DatabaseName)
	LocalAddr = fmt.Sprintf("localhost:%s", Port)
}

func getPgConnUrl(usr, passwd, host, port, database string) string {
	builder := strings.Builder{}

	_, _ = builder.WriteString("postgres://")
	_, _ = builder.WriteString(usr)

	if passwd != "" {
		_, _ = builder.WriteRune(':')
		_, _ = builder.WriteString(passwd)
	}

	_, _ = builder.WriteRune('@')
	_, _ = builder.WriteString(host)

	if port != "" {
		_, _ = builder.WriteRune(':')
		_, _ = builder.WriteString(port)
	}

	_, _ = builder.WriteRune('/')
	_, _ = builder.WriteString(database)
	_, _ = builder.WriteString("?sslmode=disable")

	return builder.String()
}
