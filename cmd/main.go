package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"helloword/internal/di"
	"go-common/library/conf/paladin.v2"
	"go-common/library/net/trace"
	"go-common/library/log"

	_ "go.uber.org/automaxprocs"
)

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("helloword start")
	trace.Init(nil)
	defer trace.Close()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("helloword exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
