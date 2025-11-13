package migrate

import (
	"context"
	"fmt"
	migrate "wallet/internal/repository/schema"

	"github.com/urfave/cli/v3"
)

var Cmd = cli.Command{
	Name:  "migrate",
	Usage: "set actual version migration",
	Flags: cmdFlags,
	OnUsageError: func(_ctx context.Context, _cmd *cli.Command, err error, _isSubcommand bool) error {
		return err
	},
	Action: run,
}

func run(ctx context.Context, cmd *cli.Command) error {
	return migrate.Up(
		ctx,
		fmt.Sprintf(
			"yugabytedb://%s:%s@%s/%s",
			cmd.String("yugabytedb-user"),
			cmd.String("yugabytedb-password"),
			cmd.String("yugabytedb-host"),
			cmd.String("yugabytedb-db-name"),
		),
	)
}
