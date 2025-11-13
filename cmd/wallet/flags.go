package wallet

import (
	"time"

	"github.com/urfave/cli/v3"
)

var cmdFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "debug",
		Usage:   "Debug",
		Sources: cli.EnvVars("DEBUG"),
		Value:   true,
	},
	&cli.StringFlag{
		Name:    "server-host",
		Usage:   "server host",
		Sources: cli.EnvVars("SERVER_HOST"),
		Value:   "localhost:3000",
	},
	&cli.BoolFlag{
		Name:    "server-debug",
		Usage:   "server debug",
		Sources: cli.EnvVars("SERVER_DEBUG"),
		Value:   true,
	},
	&cli.BoolFlag{
		Name:    "server-profiling",
		Usage:   "server profiling",
		Sources: cli.EnvVars("SERVER_PROFILING"),
		Value:   false,
	},
	&cli.IntFlag{
		Name:    "metrics-port",
		Usage:   "metrics port",
		Sources: cli.EnvVars("METRICS_PORT"),
		Value:   4000,
	},
	&cli.StringFlag{
		Name:    "yugabytedb-user",
		Usage:   "yugabytedb user",
		Sources: cli.EnvVars("YUGABYTEDB_USER"),
		Value:   "postgres",
	},
	&cli.StringFlag{
		Name:    "yugabytedb-password",
		Usage:   "yugabytedb password",
		Sources: cli.EnvVars("YUGABYTEDB_PASSWORD"),
		Value:   "postgres",
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
	// tokens
	&cli.DurationFlag{
		Name:    "access-token-lifespan",
		Usage:   "access token lifespan",
		Sources: cli.EnvVars("ACCESS_TOKEN_LIFESPAN"),
		Value:   1 * time.Hour,
	},
	&cli.StringFlag{
		Name:    "token-secret-key",
		Usage:   "token secret key",
		Sources: cli.EnvVars("TOKEN_SECRET_KEY"),
	},
	// password
	&cli.IntFlag{
		Name:    "password-hash-algorithm-id",
		Usage:   "password hash algorithm ID",
		Sources: cli.EnvVars("PASSWORD_HASH_ALGORITHM_ID"),
		Value:   255,
	},
	&cli.StringFlag{
		Name:    "password-hash-pepper",
		Usage:   "password hash pepper",
		Sources: cli.EnvVars("PASSWORD_HASH_PEPPER"),
	},
}
