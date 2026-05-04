package config

import "github.com/urfave/cli/v3"

type db struct {
	Address string
}

func (db *db) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "db.address",
			Value:       "",
			Usage:       "DB Address",
			Sources:     cli.EnvVars("DB_ADDRESS"),
			Destination: &db.Address,
		},
	}
}
