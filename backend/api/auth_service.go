package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"main/backend/config"
	"main/backend/store"
	v1 "main/proto/generated-go/api/v1"
	v1pb "main/proto/generated-go/api/v1"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// return hashed password and salt.
func hashPassword(password string, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt))
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return fmt.Sprintf("%x", hash)
}

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, r *v1pb.LoginRequest) (*v1pb.LoginResponse, error) {
	// Verify that the username and password match
	salt, passwdHash, id, err := store.GetUser(r.Name)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	hashString := hashPassword(r.Passwd, salt)
	if passwdHash != hashString {
		return &v1pb.LoginResponse{}, errors.New("Wrong username or password")
	}
	// The username and password are correct
	claims := jwt.MapClaims{
		"name": r.Name,
		"exp":  config.TokenExpiration,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(config.PrivateKey)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	sessionId := uuid.New()
	key := id
	values := map[string]string{
		"Token":     tokenString,
		"sessionId": sessionId.String(),
	}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	err = store.RdbSetSession(fmt.Sprint(key), jsonValue, config.TokenExpiration)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	return &v1pb.LoginResponse{
		Token:     tokenString,
		SessionId: sessionId.String(),
	}, nil
}

func (s *AuthService) Register(ctx context.Context, r *v1pb.RegisterRequest) (*v1pb.RegisterResponse, error) {
	// check the name if exists
	_, _, _, err := store.GetUser(r.Name)
	if err == sql.ErrNoRows {
		salt := make([]byte, 16)
		_, _ = rand.Read(salt)
		saltString := fmt.Sprintf("%x", salt)
		passwordHash := hashPassword(r.Passwd, saltString)
		err = store.CreateUser(r.Name, passwordHash, saltString, r.Email, r.Phone)
		if err != nil {
			return nil, errors.Wrap(err, "failed to register")
		}
		return &v1pb.RegisterResponse{}, nil
	} else if err == nil {
		return &v1.RegisterResponse{}, errors.New("The username already exists")
	} else {
		return &v1.RegisterResponse{}, err
	}

}
