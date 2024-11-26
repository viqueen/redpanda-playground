package main

import (
	"github.com/spf13/cobra"
	"github.com/viqueen/go-redpanda/internal/spicedb"
	"os"
)

var rootCmd = &cobra.Command{}

var migrateSchema = &cobra.Command{
	Use:   "write-schema",
	Short: "Write spicedb schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		return spicedb.WriteSchema()
	},
}

func init() {
	rootCmd.AddCommand(migrateSchema)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
