package main

import (
	"github.com/kirqz23/smart-guardian/http"
	"github.com/kirqz23/smart-guardian/pkg/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func serveCmd(cfg config.Config, log *zap.SugaredLogger) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			router := http.NewRouter(cfg, log)
			addr := ":8080"
			log.Infof("Starting server on %s", addr)
			return http.ListenAndServe(addr, router)
		},
	}
}
