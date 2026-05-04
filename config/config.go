package config

import (
	"slices"

	"github.com/urfave/cli/v3"
)

// Config -.
type Config struct {
	APP  app
	HTTP http
	JWT  jwt
	DB   db
	Log  log
}

// Flags -.
func (c *Config) Flags() []cli.Flag {
	return slices.Concat(
		c.APP.flags(),
		c.HTTP.flags(),
		c.JWT.flags(),
		c.DB.flags(),
		c.Log.flags(),
	)
}
