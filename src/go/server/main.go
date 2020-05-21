package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"server/internal"
	"syscall"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func run(args []string) error {
	server, err := internal.BuildServer(args)
	if err != nil {
		return err
	}
	setupShutdown(server)
	log.Printf("Starting server")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	return server.Serve(listener)
}

func setupShutdown(server *grpc.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("Shutting down server")
		server.GracefulStop()
	}()
}
