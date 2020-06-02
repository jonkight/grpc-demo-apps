package test

import (
	"context"
	"log"
	"net"
	"server/services"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"

	api "github.com/jonkight/grpc-demo-clients/gen/go/demo/user/v1"
)

const (
	port    = ":8080"
	address = "localhost:8080"
	bufSize = 1024 * 1024
)

var bufListener *bufconn.Listener

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}

func initConn() {
	bufListener = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	services.RegisterUserAPIv1(server)
	reflection.Register(server)
	go func() {
		if err := server.Serve(bufListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func TestUserAPIv1_CreateUser(t *testing.T) {
	ctx := context.Background()
	initConn()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(getBufDialer(bufListener)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := api.NewUserAPIClient(conn)

	userId := "12345"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.CreateUser(ctx, &api.CreateUserRequest{User: &api.User{UserId: userId}})
	if err != nil {
		t.Errorf("Could not create user: %v", err)
	}
}
