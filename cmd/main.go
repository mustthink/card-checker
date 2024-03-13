package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mustthink/card-checker/internal"
)

func main() {
	server := internal.NewServer()

	go func() {
		err := server.Run()
		if err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	err := server.Stop()
	if err != nil {
		panic(err)
	}
}
