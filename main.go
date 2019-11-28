package main

import (
	srv "github.com/securenative/{{{ .PackageName }}}/cmd/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := srv.ParseConfig()
	log.Printf("Starting {{{ .Name }}} Server...")

	shutdownHook()
	module := srv.NewModule(cfg)
	log.Printf("{{{ .Name }}} module was initiated successfully.")

	err := module.GrpcServer.Start()
	if err != nil {
		log.Printf("Failed to start the grpc server, %s", err.Error())
		os.Exit(1)
	}
}

func shutdownHook() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		recSignal := <-signalChannel
		log.Printf("System signal received (%s), exiting now...", recSignal.String())
		os.Exit(1)
	}()
}
