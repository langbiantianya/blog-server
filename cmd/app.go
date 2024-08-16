package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
		log.Println(ctx, "接收到信号", sig.String())
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
	log.Println(ctx, "服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
