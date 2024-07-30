package store

import (
	"time"
)

func CreateUser(name, hashedPasswd, salt, email, phone string) error {
	sql := `INSERT INTO users (name, passwd_hash, salt, email, phone)
VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sql, name, hashedPasswd, salt, email, phone)
	if err != nil {
		return err
	}
	return nil
}
func GetUser(name string) (string, string, int, error) {
	query := "SELECT salt, passwd_hash,id FROM users WHERE name = $1"
	var salt, passwdHash string
	var id int
	err := db.QueryRow(query, name).Scan(&salt, &passwdHash, &id)
	return salt, passwdHash, id, err
}
func RdbSetSession(key string, jsonValue []byte, expiration time.Duration) error {
	err := rdb.Set(ctx, string(key), jsonValue, expiration).Err()
	return err
}
func RdbDelSession(key string) error {
	return rdb.Del(ctx, key).Err()
}
