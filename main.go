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
				go cmd.StartStaticServer()
				go cmd.StartApiServer()
				return nil
			}, func() error {

				return nil
			})
			return nil
		},
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "path",
				DefaultText: "./www/html",
				Usage:       "静态页面文件夹",
				Required:    false,
			},
			&cli.PathFlag{
				Name:        "md",
				DefaultText: "./md",
				Usage:       "Markdown存放路径",
				Required:    false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "port",
				Usage: "端口",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "static",
						DefaultText: "8000",
						Usage:       "静态页面端口",
						Required:    false,
					}, &cli.IntFlag{
						Name:        "api",
						DefaultText: "8001",
						Usage:       "控制台端口",
						Required:    false,
					},
				},
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
