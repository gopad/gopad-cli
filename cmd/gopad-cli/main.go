package main

import (
	"os"

	"github.com/gopad/gopad-cli/pkg/command"
	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("GOPAD_CLI_ENV_FILE"); env != "" {
		_ = godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
