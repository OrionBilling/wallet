package api

import (
	"wallet/api/grpc"
)

type Container struct {
	GRPCServer *grpc.Resolver
}

// NewContainer creates a new application struct
func NewContainer(
	server *grpc.Resolver,
) Container {
	return Container{
		GRPCServer: server,
	}
}
