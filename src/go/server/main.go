package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"server/services"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func run(args []string) error {
	server, err := buildServer(args)
	if err != nil {
		return err
	}
	setupShutdown(server)
	log.Printf("Starting server")
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		return err
	}
	return server.Serve(listener)
}

func buildServer(args []string) (*grpc.Server, error) {
	if len(args) < 1 {
		return nil, errors.New("Server requires valid services specified to run.")
	}
	server := grpc.NewServer()
	for _, service := range args {
		switch service {
		case "userV1":
			services.RegisterUserAPIv1(server)
		case "leaderboardV1":
		case "notificationV1":
			return nil, errors.New(fmt.Sprintf("Service '%s' is not implemented.", service))
		default:
			return nil, errors.New(fmt.Sprintf("Invalid service name '%s'.", service))
		}
	}
	return server, nil
}

func setupShutdown(server *grpc.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("Shutting down server")
		server.GracefulStop()
		os.Exit(0)
	}()
}
