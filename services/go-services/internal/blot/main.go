package main

import (
	blotservicepb "blot/internal/blot/gen/blotservice/v1beta1"
	"blot/internal/blot/ports"
	"blot/internal/blot/service"
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
