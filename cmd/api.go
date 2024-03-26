package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sznborges/to_do_list/infra/server"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Initializes the codebase running as HTTP API",
		Long:  "Initializes the codebase running as HTTP API",
		RunE:  apiExecute,
	}
)

func init() {
	rootCmd.AddCommand(apiCommand)
}

func apiExecute(cmd *cobra.Command, args []string) error {
	server.StartHTTP()
	return nil
}