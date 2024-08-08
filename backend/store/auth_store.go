package store

import (
	"context"
	"database/sql"
	"fmt"
	"main/backend/config"

	"github.com/pkg/errors"
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
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
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

func SetSession(ctx context.Context, key string, jsonValue []byte) error {
	if err := rdb.Set(ctx, redisKey(string(redisSession), key), jsonValue, config.AccessTokenExpiration).Err(); err != nil {
		return errors.Wrapf(err, "failed to set session")
	}
	return nil
}

func DelSession(ctx context.Context, session string) error {
	if err := rdb.Del(ctx, redisKey(string(redisSession), session)).Err(); err != nil {
		return errors.Wrapf(err, fmt.Sprintf("failed to delete session %q", session))
	}
	return nil
}

func DeactivateAccessToken(ctx context.Context, token string) error {
	if err := rdb.SAdd(ctx, redisKey(string(redisExpiredAccessToken)), token).Err(); err != nil {
		return errors.Wrapf(err, fmt.Sprintf("failed to deactivate access token %q", token))
	}
	return nil
}
