package migrate

import "github.com/urfave/cli/v3"

var cmdFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "yugabytedb-user",
		Usage:   "yugabytedb user",
		Sources: cli.EnvVars("YUGABYTEDB_USER"),
		Value:   "admin",
	},
	&cli.StringFlag{
		Name:    "yugabytedb-password",
		Usage:   "yugabytedb password",
		Sources: cli.EnvVars("YUGABYTEDB_PASSWORD"),
		Value:   "admin",
	},
	&cli.StringFlag{
		Name:    "yugabytedb-host",
		Usage:   "yugabytedb host",
		Sources: cli.EnvVars("YUGABYTEDB_HOST"),
		Value:   "localhost:5432",
	},
	&cli.StringFlag{
		Name:    "yugabytedb-db-name",
		Usage:   "yugabytedb database name",
		Sources: cli.EnvVars("YUGABYTEDB_DB_NAME"),
		Value:   "sn",
	},
	&cli.BoolFlag{
		Name:    "yugabytedb-disable-tls",
		Usage:   "yugabytedb disable tls",
		Sources: cli.EnvVars("YUGABYTEDB_DISABLE_TLS"),
		Value:   true,
	},
}
