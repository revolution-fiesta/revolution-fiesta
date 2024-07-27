package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"main/backend/store"
	v1pb "main/proto/generated-go/api/v1"

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
	return &v1pb.LoginResponse{
		Token: "I am token...",
	}, nil
}

func (s *AuthService) Register(ctx context.Context, r *v1pb.RegisterRequest) (*v1pb.RegisterResponse, error) {
	passwordHash, salt := hashPassword(r.Passwd)
	err := store.CreateUser(r.Name, passwordHash, salt, r.Email, r.Phone)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register")
	}
	return &v1pb.RegisterResponse{}, nil
}
