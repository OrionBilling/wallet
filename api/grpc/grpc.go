package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"wallet/api/grpc/reserve"
	"wallet/internal/core"
	pb "wallet/pkg/grpc/reserve/generated"
	"wallet/pkg/logger"
)

type (
	Resolver struct {
		grpcServer *grpc.Server
		config     Config
		// services
		reserveService pb.WalletReservationServiceServer
	}

	Config struct {
		Host string
	}
)

func NewResolver(
	config Config,
	wallet core.WalletService,
) *Resolver {
	return &Resolver{
		grpcServer:     grpc.NewServer(),
		config:         config,
		reserveService: reserve.NewService(wallet),
	}
}

func (r *Resolver) Run(ctx context.Context) {
	logger.Log().Info(ctx, "Run gRPC server")

	lis, err := net.Listen("tcp", r.config.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r.init()

	if err := r.grpcServer.Serve(lis); err != nil {
		logger.Log().Fatalf(ctx, "failed to listen to socket:")
	}
}

func (r *Resolver) init() {
	pb.RegisterWalletReservationServiceServer(r.grpcServer, r.reserveService)
}

func (r *Resolver) Shutdown(ctx context.Context) error {
	r.grpcServer.Stop()

	logger.Log().Info(ctx, "gRPC sever stopped")

	return nil
}
