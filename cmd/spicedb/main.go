package main

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/viqueen/go-redpanda/internal/spicedb"
	"os"
)

var rootCmd = &cobra.Command{}

var migrateSchema = &cobra.Command{
	Use:   "write-schema",
	Short: "Write spicedb schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		return spicedb.WriteSchema(context.Background())
	},
}

var setupData = &cobra.Command{
	Use:   "setup-data",
	Short: "Setup spicedb data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return spicedb.SetupData(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(migrateSchema)
	rootCmd.AddCommand(setupData)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
