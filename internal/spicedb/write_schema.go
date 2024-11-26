package spicedb

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"log"
	"os"
	"path/filepath"
)

func WriteSchema() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	schemaFile := filepath.Join(wd, "internal", "spicedb", "schema.zed")
	schemaContent, err := os.ReadFile(schemaFile)
	if err != nil {
		return err
	}

	return withClient(func(client *authzed.Client) error {
		schema := string(schemaContent)
		request := &pb.WriteSchemaRequest{
			Schema: schema,
		}
		response, writeSchemaErr := client.WriteSchema(context.Background(), request)
		if writeSchemaErr != nil {
			return writeSchemaErr
		}
		log.Printf("Schema written: token=%s", response.GetWrittenAt())
		return nil
	})
}
