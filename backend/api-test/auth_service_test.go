package apitest

import (
	"context"
	"main/backend/config"
	v1pb "main/proto/generated-go/api/v1"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (v1pb.AuthServiceClient, error) {
	client, err := grpc.NewClient(config.LocalAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return v1pb.NewAuthServiceClient(client), nil
}

func TestRegister(t *testing.T) {
	t.Skip()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	client.Register(context.Background(), &v1pb.RegisterRequest{
		Name:   "",
		Passwd: "abc",
	})
}
