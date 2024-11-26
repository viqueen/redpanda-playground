package spicedb

import (
	"context"
	"errors"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	spicedbEndpoint = "localhost:50051"
)

func withClient(ctx context.Context, handler func(ctx context.Context, client *authzed.Client) error) error {
	client, err := authzed.NewClient(spicedbEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New("failed to create client")
	}
	return handler(ctx, client)
}
