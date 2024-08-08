package apitest

import (
	"context"
	"fmt"
	"main/backend/api/utils"
	"main/backend/config"
	v1pb "main/proto/generated-go/api/v1"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
	resp, err := client.Register(context.Background(), &v1pb.RegisterRequest{
		Name:   "azusaings@gmail.com",
		Passwd: "Aa020111",
	})
	fmt.Println(resp, err)
}

func TestLogin(t *testing.T) {
	t.Skip()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	resp, err := client.Login(ctx, &v1pb.LoginRequest{
		Name:   "azusa@xxx.com",
		Passwd: "xxxxxx",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}

func TestLogout(t *testing.T) {
	t.Skip()
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	resp, err := client.Login(ctx, &v1pb.LoginRequest{
		Name:   "azusa@xxx.com",
		Passwd: "xxxxxx",
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(utils.HttpHeaderAuthorization, fmt.Sprintf("bearer %s", resp.Token)))
	loResp, err := client.Logout(ctx, &v1pb.LogoutRequest{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(loResp)
}
