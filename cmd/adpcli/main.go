package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = 443
	DEFAULT_USER = "adpuser"
)

func main() {

	app := &cli.App{
		Name:    "adp-cli",
		Version: "0.1-alpha",
		Usage:   "Command Line Interface to access ADP rest service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"t"},
				Usage:   "hostname which runs ADP service",
				Value:   DEFAULT_HOST,
			},
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"o"},
				Usage:   "port",
				Value:   DEFAULT_PORT,
			},
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "ADP User",
				Value:   DEFAULT_USER,
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "ADP User Password",
				EnvVars: []string{"ADP_USER_PASSWORD"},
			},
			&cli.StringFlag{
				Name:    "taskAccessKey",
				Aliases: []string{"k"},
				Usage:   "ADP Task Access Key (for Query Postgresql DB Task)",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Debug Mode: Log to stderr and AdpTask log enabled",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "pretty",
				Aliases: []string{"b"},
				Usage:   "Pretty Print",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "async",
				Aliases: []string{"a"},
				Usage:   "Asynchronous",
				Value:   false,
			},
		},
		Commands: Commands,
		Before: func(c *cli.Context) error {
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
			if c.Bool("debug") {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err)
	}
}
