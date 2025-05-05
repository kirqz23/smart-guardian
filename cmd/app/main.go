package main

import (
	"fmt"
	"os"

	"github.com/kirqz23/smart-guardian/pkg/config"
	"github.com/kirqz23/smart-guardian/pkg/logger"
	"github.com/spf13/cobra"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading config:", err)
		os.Exit(1)
	}
	log := logger.New(cfg.LogLevel)

	root := &cobra.Command{Use: "smart-guardian"}

	root.AddCommand(migrateCmd(cfg, log))
	root.AddCommand(serveCmd(cfg, log))

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
