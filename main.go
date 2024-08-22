package main

import (
	"blog-server/cmd"
	"blog-server/internal/constantx"
	"blog-server/internal/entity"

	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := &cli.App{
		Name:  "blog",
		Usage: "博客",
		Action: func(ctx *cli.Context) error {
			cmd.Run(context.Background(), func() error {
				path := ctx.String("path")
				port := ctx.Int64Slice("port")
				db, err := gorm.Open(sqlite.Open("./static/sqlite/sqlite.db"), &gorm.Config{})
				if err != nil {
					panic(err)
				}
				constantx.Db = db
				db.AutoMigrate(&entity.Essay{}, &entity.Tag{})
				go cmd.StartStaticServer(path, int(port[0]))
				go cmd.StartApiServer(int(port[1]), InitApis())
				return nil
			}, func() error {

				return nil
			})
			return nil
		},
		Flags: []cli.Flag{
			&cli.Int64SliceFlag{
				Name:        "port",
				DefaultText: "8000,8001",
				Value:       cli.NewInt64Slice(8000, 8001),
				Usage:       "端口,第一个是静态代理端口，第二个是后台api",
				Required:    false,
			},
			&cli.PathFlag{
				Name:        "path",
				DefaultText: "./www/html",
				Value:       "./www/html",
				Usage:       "静态页面文件夹",
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
