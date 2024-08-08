package main

import (
	"context"
	"fmt"
	"log/slog"
	"main/backend/api"
	"main/backend/api/interceptors"
	"main/backend/api/utils"
	"main/backend/config"
	"main/backend/store"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := store.Init()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to initiate store layer: %s", err.Error()))
		return
	}

	err = store.Migrate()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to migrate: %s", err.Error()))
		return
	}

	config.PrivateKey, err = utils.GeneratePrivateKey()
	if err != nil {
		slog.Error("failed to generate rsa key pairs")
		return
	}

	// config grpc server.
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AuthInterceptor),
	)
	api.ConfigRouter(server)
	reflection.Register(server)

	// signal handler.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		postgreErr, redisErr := store.Close()
		if redisErr != nil {
			slog.Warn(err.Error())
		}
		if postgreErr != nil {
			slog.Warn(err.Error())
		}
		slog.Info(fmt.Sprintf("%s received, bye~", sig.String()))
		cancel()
	}()

	// start grpc server.
	go func() {
		listener, err := net.Listen("tcp", ":"+config.Port)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to listen at port %s", config.Port))
			return
		}
		if err := server.Serve(listener); err != nil {
			slog.Error("failed to start service")
		}
		cancel()
	}()

	<-ctx.Done()
}
