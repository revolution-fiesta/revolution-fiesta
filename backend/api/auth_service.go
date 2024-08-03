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
	user, err := store.GetUserByName(r.Name)
	if err != nil {
		return &v1pb.LoginResponse{}, err
	}
	hashString := hashPassword(r.Passwd, user.Salt)
	if user.PasswdHash != hashString {
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
	key := user.Id
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
	// check if the user exists.
	_, err := store.GetUserByName(r.Name)
	if err == nil {
		return nil, errors.New("The username already exists")
	} else if err == sql.ErrNoRows {
		salt := make([]byte, 16)
		_, _ = rand.Read(salt)
		saltString := fmt.Sprintf("%x", salt)
		passwordHash := hashPassword(r.Passwd, saltString)
		err = store.CreateUser(r.Name, passwordHash, saltString, r.Email, r.Phone, store.UserTypeRegular)
		if err != nil {
			return nil, errors.Wrap(err, "failed to register")
		}
	} else {
		return nil, err
	}
	return &v1pb.RegisterResponse{}, nil
}

func (s *AuthService) Logout(ctx context.Context, r *v1pb.LogoutRequest) (*v1pb.LogoutResponse, error) {
	return &v1pb.LogoutResponse{}, store.RdbDelSession(fmt.Sprint(r.Id))
}
