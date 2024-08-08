package interceptors

import (
	"context"
	"fmt"
	"main/backend/api/utils"
	"main/backend/config"
	v1 "main/proto/generated-go/api/v1"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// TODO(tommy): expired tokens...
	if requireAuth(info.FullMethod) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("failed to get auth metadata")
		}

		vals := meta.Get(utils.HttpHeaderAuthorization)
		if len(vals) != 1 {
			return nil, errors.New("duplicate authorization values")
		}
		tokens := strings.Split(vals[0], " ")
		if len(tokens) != 2 || strings.ToLower(tokens[0]) != utils.HttpHeaderBearer {
			return nil, errors.New(fmt.Sprintf("invalid access token format: %v", tokens))
		}

		userId, err := utils.ValidateAccessToken(tokens[1], &config.PrivateKey.PublicKey)
		if err != nil {
			return nil, errors.Wrap(err, "failed to validate access token")
		}
		utils.WithUserId(&ctx, userId)
	}
	return handler(ctx, req)
}

func requireAuth(methodName string) bool {
	switch methodName {
	case v1.AuthService_Register_FullMethodName,
		v1.AuthService_Login_FullMethodName:
		return false
	default:
	}
	return true
}
