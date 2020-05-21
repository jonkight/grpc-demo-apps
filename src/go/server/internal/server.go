package internal

import (
	"errors"
	"fmt"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BuildServer(args []string) (*grpc.Server, error) {
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
	reflection.Register(server)
	return server, nil
}
