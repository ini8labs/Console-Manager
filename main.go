package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/ini8labs/console-manager/src/server"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	go func() {
		if err := server.NewServer(":8080", logger); err != nil {
			panic(err.Error())
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-interrupt

	logger.Info("Closing the Server")
}
