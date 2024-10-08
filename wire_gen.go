// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"blog-server/cmd"
	"blog-server/internal/repo"
	"blog-server/internal/routes"
	"blog-server/internal/service"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitApp(db *gorm.DB, ctx *cli.Context) cmd.App {
	iEssayRepo := repo.NewEssayRepo(db)
	iTagRepo := repo.NewTagRepo(db)
	iEssayService := service.NewEssayService(iEssayRepo, iTagRepo)
	iTagService := service.NewTagService(iTagRepo)
	iEssayRouter := routes.NewEssayRouter(iEssayService, iTagService)
	iTagRouter := routes.NewTagRouter(iTagService)
	apiRoutes := cmd.NewApiRoutes(iEssayRouter, iTagRouter)
	app := cmd.NewApp(apiRoutes)
	return app
}
