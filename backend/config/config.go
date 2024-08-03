package config

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"strings"
	"time"
)

const (
	Version         string        = "0.0.1"
	Port            string        = "8080"
	DatabaseUsr     string        = "postgres"
	DatabasePasswd  string        = "postgres"
	DatabaseHost    string        = "114.132.70.108"
	DatabasePort    string        = "5432"
	DatabaseName    string        = "postgres"
	SchemaFilePath  string        = "../schema.sql"
	RedisAddr       string        = "114.132.70.108:6379"
	RedisPassword   string        = ""
	RedisDB                       = 0
	TokenExpiration time.Duration = time.Hour / 4
)

var (
	DatabaseUrl string
	LocalAddr   string
	PublicKey   *rsa.PublicKey
	PrivateKey  *rsa.PrivateKey
)

func createKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func init() {
	DatabaseUrl = getPgConnUrl(DatabaseUsr, DatabasePasswd, DatabaseHost, DatabasePort, DatabaseName)
	PrivateKey, PublicKey, _ = createKey()
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
