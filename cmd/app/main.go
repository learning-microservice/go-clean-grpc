package main

import (
	"context"
	"log"
	"os"

	"go-clean-grpc/config"
	"go-clean-grpc/internal/app"

	"github.com/urfave/cli/v3"
)

var (
	name    = "go-clean-grpc"
	version = "v1.0.0-alpha.1"
)

func main() {
	// 1. 引数があればそれを使用
	// 2. 引数がなく環境変数があればそれを使用
	// 3. どちらもなければ DefaultValue を使用
	var cfg config.Config

	cmd := cli.Command{
		Name:    name,
		Version: version,
		Flags:   cfg.Flags(),
		Action: func(context.Context, *cli.Command) error {
			// bind app parameters
			cfg.APP.Name = name
			cfg.APP.Version = version

			return app.Run(&cfg)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
