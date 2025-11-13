package inject

import (
	grpcserver "wallet/api/grpc"

	"github.com/google/wire"
	"github.com/urfave/cli/v3"
)

// wire Set for loading the server.
var serverSet = wire.NewSet( // nolint: unused // it used for wire autogeneration
	grpcserver.NewResolver,
	provideGRPCServerConfig,
)

func provideGRPCServerConfig(c *cli.Command) grpcserver.Config {
	return grpcserver.Config{
		Host: c.String("grpc-host"),
	}
}
