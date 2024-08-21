//go:build wireinject
// +build wireinject

package main

import (
	"blog-server/cmd"
	"blog-server/internal/repo"
	"blog-server/internal/routes"
	"blog-server/internal/service"

	"github.com/google/wire"
)

func InitApis() *cmd.ApiRoutes {
	wire.Build(
		repo.Set,
		service.Set,
		routes.Set,
		cmd.Set,
	)
	return nil
}
