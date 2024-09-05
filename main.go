package main

import (
	"blog-server/cmd"

	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "blog",
		Usage: "博客",
		Action: func(ctx *cli.Context) error {
			cmd.Run(context.Background(), func() error {
				InitApp(ctx).StartServer()
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
