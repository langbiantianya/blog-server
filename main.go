package main

import (
	"blog-server/cmd"
	"blog-server/internal/conf"
	"blog-server/internal/entity"
	"path"

	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := &cli.App{
		Name:  "blog",
		Usage: "博客",
		Action: func(ctx *cli.Context) error {
			cmd.Run(context.Background(), func() error {
				conf.InitConfig(ctx)
				err := os.MkdirAll(path.Clean("./static/sqlite/"), 0755)
				if err != nil {
					panic(err)
				}
				db, err := gorm.Open(sqlite.Open("./static/sqlite/sqlite.db"), &gorm.Config{
					Logger: logger.Default.LogMode(logger.Info),
				})
				if err != nil {
					panic(err)
				}
				db.AutoMigrate(entity.MigeateEntity...)
				InitApp(db, ctx).StartServer()
				return nil
			}, func() error {

				return nil
			})
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntSliceFlag{
				Name:        "port",
				DefaultText: "8000,8001",
				Value:       cli.NewIntSlice(8000, 8001),
				Usage:       "端口,第一个是静态代理端口，第二个是后台api",
				Required:    false,
			},
			&cli.PathFlag{
				Name:        "staticPath",
				DefaultText: "./static",
				Value:       "./static",
				Usage:       "存放静态资源的路径",
				Required:    false,
			},
			&cli.PathFlag{
				Name:        "outPath",
				DefaultText: "./www/blog",
				Value:       "./www/blog",
				Usage:       "生成的静态页面根路径",
				Required:    false,
			},
			&cli.PathFlag{
				Name:        "md",
				DefaultText: "./md",
				Value:       "./md",
				Usage:       "Markdown存放路径",
				Required:    false,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
