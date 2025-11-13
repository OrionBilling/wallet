package main

import (
	"context"
	"os"

	cli "github.com/urfave/cli/v3"

	"wallet/cmd/migrate"
	wallet "wallet/cmd/wallet"
	"wallet/pkg/logger"
	"wallet/version"
)

// @Version 0.0.0
// @Title Wallet
// @Description Orion Billing Wallet backend
func main() {
	var globalFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "Debug",
			Sources: cli.EnvVars("DEBUG"),
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "disable-stack-trace",
			Usage:   "Disable stack trace",
			Sources: cli.EnvVars("DISABLE_STACK_TRACE"),
			Value:   true,
		},
	}

	app := &cli.Command{
		Usage: "Wallet",
		Commands: []*cli.Command{
			&wallet.Cmd,
			&migrate.Cmd,
		},
		EnableShellCompletion: true,
		Flags:                 globalFlags,
		Version:               version.Version + " (" + version.GitCommit + ")",
		OnUsageError: func(_ context.Context, _ *cli.Command, err error, _ bool) error {
			return err
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			serviceName := "wallet"

			if cmd.Bool("debug") {
				logger.SetDebugLogger(serviceName, cmd.Bool("disable-stack-trace"))
			} else {
				logger.SetProductionLogger(serviceName)
			}

			return ctx, nil
		},
	}

	appCtx := context.Background()
	if err := app.Run(appCtx, os.Args); err != nil {
		logger.Log().Errorf(appCtx, "application was stopped: %s", err)
	}
}
