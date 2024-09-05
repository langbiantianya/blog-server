package cmd

import (
	"blog-server/internal/constantx"
	"blog-server/internal/entity"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	staticPath string
	apiPort    int
	staticPort int
	apiRoutes  *ApiRoutes
}

func NewConfig(ctx *cli.Context, apiRoutes *ApiRoutes) Config {
	port := ctx.IntSlice("port")
	return Config{
		staticPath: ctx.String("path"),
		apiPort:    port[1],
		staticPort: port[0],
		apiRoutes:  apiRoutes,
	}
}

func (c Config) StartServer() {
	db, err := gorm.Open(sqlite.Open("./static/sqlite/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	constantx.Db = db
	db.AutoMigrate(&entity.Essay{}, &entity.Tag{})
	go StartStaticServer(c.staticPath, c.staticPort)
	go StartApiServer(c.apiPort, c.apiRoutes)
}

func Run(ctx context.Context, start func() error, clean func() error) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	err := start()
	if err != nil {
		log.Fatalln(err)
	}
EXIT:
	for {
		sig := <-sc
		log.Print(ctx, "接收到信号", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	err = clean()
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(ctx, "服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
