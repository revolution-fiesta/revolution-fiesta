package api

import (
	"context"
	v1pb "main/proto/generated-go/api/v1"
)

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Login(ctx context.Context, r *v1pb.LoginRequest) (*v1pb.LoginResponse, error) {
	return &v1pb.LoginResponse{
		Token: "I am token...",
	}, nil
}
