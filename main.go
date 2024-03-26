package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sznborges/to_do_list/cmd"
	"github.com/sznborges/to_do_list/config"
	"github.com/sznborges/to_do_list/infra/logger"
)

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		logger.Logger.Fatal(err)
	}
	env := config.GetString("ENV")
	if env != "" {
		if err := godotenv.Load(fmt.Sprintf("%s.%s", filepath.Join(currentDir, ".env"), env)); err != nil {
			logger.Logger.Fatal(err)
		}
	}
	if err := godotenv.Load(filepath.Join(currentDir, ".env")); err != nil {
		logger.Logger.Fatal(err)
	}
}
func main() {
	cmd.Execute()
}

