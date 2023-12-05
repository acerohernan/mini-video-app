package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acerohernan/mini-video-app/pkg/server"
)

func main() {
	server := server.NewServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan
		log.Print("exit requested, closing server...")
		server.Stop()
	}()

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
