package internal

import (
	"errors"
	"fmt"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// BuildServer builds the internal gRPC server based on the passed arguments.
func BuildServer(args []string) (*grpc.Server, error) {
	if len(args) < 1 {
		return nil, errors.New("server requires valid services specified to run")
	}
	server := grpc.NewServer()
	for _, service := range args {
		switch service {
		case "userV1":
			services.RegisterUserAPIv1(server)
		case "leaderboardV1":
		case "notificationV1":
			return nil, fmt.Errorf("service '%s' is not implemented", service)
		default:
			return nil, fmt.Errorf("invalid service name '%s'", service)
		}
	}
	reflection.Register(server)
	return server, nil
}
