package config

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v3"
)

type jwt struct {
	Secret   string
	Duration time.Duration
}

func (jwt *jwt) flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "jwt.secret",
			Value:       "",
			Usage:       "JWT Token Secret",
			Sources:     cli.EnvVars("JWT_SECRET"),
			Destination: &jwt.Secret,
			Required:    true,
		},
		&cli.DurationFlag{
			Name:        "jwt.duration",
			Value:       24 * time.Hour,
			Usage:       "JWT Token Duration",
			Sources:     cli.EnvVars("JWT_DURATION"),
			Destination: &jwt.Duration,
			Validator: func(d time.Duration) error {
				if d <= 0 {
					return fmt.Errorf("")
				}
				return nil
			},
		},
	}
}
