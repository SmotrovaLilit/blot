package main

import (
	"blot/internal/blot/ports"
	"blot/internal/blot/service"
	blotservicepb "blot/internal/common/gen-proto/blotservice/v1beta1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"blot/internal/common/server"
)

func main() {
	ctx := context.Background()
	app := service.NewApplication(ctx)

	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(app)
		reflection.Register(server)
		blotservicepb.RegisterBlotServiceServer(server, svc)
	})
}
