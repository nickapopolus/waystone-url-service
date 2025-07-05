package main

import (
	"fmt"
	grpcCon "github.com/nickapopolus/waystone-url-service/internal/grpc"
	urlgrpcv1 "github.com/nickapopolus/waystone-url-service/proto/urlservice/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := &Config{}
	app.URLGRPC = grpcCon.NewURLServiceGRPC()

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50001"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		//TODO: Log this to the logger
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	urlgrpcv1.RegisterURLServiceServer(grpcServer, app.URLGRPC)
	go func() {
		log.Printf("Starting gRPC server on :%d", grpcPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")
	grpcServer.GracefulStop()
}
