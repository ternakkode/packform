package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ternakkode/packform-backend/internal/config"
	"github.com/ternakkode/packform-backend/internal/httpserver"
	"github.com/ternakkode/packform-backend/pkg/bunclient"
)

func main() {
	c := config.Init()
	dbClient := bunclient.InitDB(&c.DB)

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	svr := httpserver.InitAndStart(c.APPName, c.APPPort)

	<-sc

	svr.Shutdown(context.Background())
	dbClient.Close()
}
