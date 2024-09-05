//go:build wireinject
// +build wireinject

package main

import (
	"blog-server/cmd"
	"blog-server/internal/repo"
	"blog-server/internal/routes"
	"blog-server/internal/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

func InitApp(ctx *cli.Context) cmd.Config {
	wire.Build(
		repo.Set,
		service.Set,
		routes.Set,
		cmd.Set,
	)
	return cmd.Config{}
}
