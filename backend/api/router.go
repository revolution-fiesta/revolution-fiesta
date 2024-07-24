package api

import (
	v1pb "main/proto/generated-go/api/v1"

	"google.golang.org/grpc"
)

func ConfigRouter(s *grpc.Server) {
	v1pb.RegisterAuthServiceServer(s, &AuthService{})
}
