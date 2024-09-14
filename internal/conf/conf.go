package conf

import "github.com/urfave/cli/v2"

type Config struct {
	StaticPath string
	ApiPort    int
	StaticPort int
}

var conf Config

func GetConfig() Config {
	return conf
}

func InitConfig(ctx *cli.Context) {
	port := ctx.IntSlice("port")
	conf = Config{
		StaticPath: ctx.String("path"),
		ApiPort:    port[1],
		StaticPort: port[0],
	}
}
