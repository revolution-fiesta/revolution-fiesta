package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"main/backend/config"
	"main/backend/store"
	v1 "main/proto/generated-go/api/v1"
	v1pb "main/proto/generated-go/api/v1"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// return hashed password and salt.
func hashPassword(password string) (string, string) {
	randBytes := make([]byte, 16)
	_, _ = rand.Read(randBytes)

	hasher := sha256.New()
	hasher.Write(randBytes)
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)

	return fmt.Sprintf("%x", hash), fmt.Sprintf("%x", randBytes)
}

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, r *v1pb.LoginRequest) (*v1pb.LoginResponse, error) {
	// Verify that the username and password match
	salt, passwd_hash, id, err := store.GetUser(r.Name)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	saltByte := []byte(salt)
	hasher := sha256.New()
	hasher.Write(saltByte)
	hasher.Write([]byte(r.Passwd))
	hash := hasher.Sum(nil)
	if passwd_hash != string(hash) {
		return &v1pb.LoginResponse{}, errors.New("Wrong username or password")
	}
	// The username and password are correct
	claims := jwt.MapClaims{
		"name": r.Name,
		"exp":  time.Now().Add(time.Hour * 1 / 4).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(config.PrivateKey)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	session_id := uuid.New()
	key := id
	value1 := tokenString
	value2 := session_id
	values := map[string]string{
		"value1": value1,
		"value2": value2.String(),
	}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	expiration := time.Hour / 4
	err = store.RdbSetx(string(key), jsonValue, expiration)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	return &v1pb.LoginResponse{
		Token:     tokenString,
		SessionId: session_id.String(),
	}, nil
}

func (s *AuthService) Register(ctx context.Context, r *v1pb.RegisterRequest) (*v1pb.RegisterResponse, error) {
	// check the name if exists
	exists, err := store.CheckNameIfExists(r.Name)
	if err != nil {
		return &v1.RegisterResponse{}, err
	}
	if exists {
		return &v1pb.RegisterResponse{}, errors.New("The username already exists")
	}
	passwordHash, salt := hashPassword(r.Passwd)
	err = store.CreateUser(r.Name, passwordHash, salt, r.Email, r.Phone)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register")
	}
	return &v1pb.RegisterResponse{}, nil
}
