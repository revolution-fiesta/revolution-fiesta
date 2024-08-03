package store

import (
	"time"
)

type UserType string

const (
	UserTypeAdmin   = "ADMIN"
	UserTypeRegular = "REGULAR"
)

type User struct {
	Id         int
	PasswdHash string
	Salt       string
	Name       string
	Email      string
	Phone      string
	Type       UserType
}

func CreateUser(name, hashedPasswd, salt, email, phone, typ string) error {
	sql := `INSERT INTO users (name, passwd_hash, salt, email, phone, type)
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sql, name, hashedPasswd, salt, email, phone, typ)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByName(name string) (*User, error) {
	query := "SELECT id, passwd_hash, salt, name, email, phone, type FROM users WHERE name = $1"
	var passwdHash, salt, usrName, email, phone, typ string
	var id int
	err := db.QueryRow(query, name).Scan(&id, &passwdHash, &salt, &usrName, &email, &phone, &typ)
	return &User{
		Id:         id,
		PasswdHash: passwdHash,
		Salt:       salt,
		Name:       usrName,
		Email:      email,
		Phone:      phone,
		Type:       UserType(typ),
	}, err
}

func RdbSetSession(key string, jsonValue []byte, expiration time.Duration) error {
	err := rdb.Set(ctx, string(key), jsonValue, expiration).Err()
	return err
}

func RdbDelSession(key string) error {
	return rdb.Del(ctx, key).Err()
}
