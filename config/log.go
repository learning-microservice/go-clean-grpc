package config

import (
	"github.com/urfave/cli/v3"
)

type log struct {
	Level string
}

func (log *log) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "DEBUG",
			Usage:       "Logging Level",
			Sources:     cli.EnvVars("LOG_LEVEL"),
			Destination: &log.Level,
		},
	}
}
