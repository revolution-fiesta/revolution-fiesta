package api

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log/slog"
	"main/backend/config"
	v1pb "main/proto/generated-go/api/v1"

	_ "github.com/lib/pq"
)

func hashPassword(password string) string {
	passwordBytes := []byte(password)
	hash := sha256.Sum256(passwordBytes)
	return fmt.Sprintf("%x", hash)
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
	var passwordHash string = hashPassword(r.Passwd)

	//connect to postgre
	db, err := sql.Open("postgres", config.DatabaseUrl)
	if err != nil {
		slog.Error("Failed to connect to postgre database")
		return nil, err
	}
	defer db.Close()
	// insert data to database
	sqlStatement := `INSERT INTO users (name, passwd_hash,email,phone,id)
		VALUES ($1, $2,$3,$4,$5)`
	_, err = db.Exec(sqlStatement, r.Name, passwordHash, r.Email, r.Phone, r.Id)
	if err != nil {
		return nil, err
	}
	return &v1pb.RegisterResponse{}, nil
}
