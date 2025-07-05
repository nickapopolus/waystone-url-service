package main

import grpcCon "github.com/nickapopolus/waystone-url-service/internal/grpc"

type Config struct {
	URLGRPC *grpcCon.URLServiceGRPC
}
