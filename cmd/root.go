package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sznborges/to_do_list/infra/logger"
)

var rootCmd = &cobra.Command{
	Use:   "painel-sbf-backend",
	Short: "This project is written in [Go](https://go.dev).",
	Long:  "This project is written in [Go](https://go.dev).",
}

// Execute root command
func Execute() {
	defer func() {
		err := recover()
		if err != nil {
			logger.Logger.Fatalf("unexpected error while executing command %v", err)
		}
	}()
	err := rootCmd.Execute()
	if err != nil {
		panic("error while executing command")
	}
}