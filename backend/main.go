package main

import (
	"context"
	"fmt"
	"log/slog"
	"main/backend/api"
	"main/backend/config"
	"main/backend/migration"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	migration.Migrate()

	// config grpc server.
	server := grpc.NewServer()
	api.ConfigRouter(server)
	reflection.Register(server)

	// signal handler.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		slog.Info(fmt.Sprintf("%s received, bye~", sig.String()))
		cancel()
	}()

	// start grpc server.
	go func() {
		listener, err := net.Listen("tcp", config.Port)
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
