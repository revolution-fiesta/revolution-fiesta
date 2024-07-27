package api

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log/slog"
	"main/backend/config"
	v1pb "main/proto/generated-go/api/v1"

	_ "github.com/lib/pq"
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

	//connect to postgre
	db, err := sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		slog.Error("Failed to connect to postgre database")
		return nil, err
	}
	defer db.Close()
	// insert data to database
	sqlStatement := `INSERT INTO users (name, passwd_hash, salt, email, phone)
		VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(sqlStatement, r.Name, passwordHash, salt, r.Email, r.Phone)
	if err != nil {
		return nil, err
	}
	return &v1pb.RegisterResponse{}, nil
}
