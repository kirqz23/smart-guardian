package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kirqz23/smart-guardian/pkg/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func migrateCmd(cfg config.Config, log *zap.SugaredLogger) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infof("Starting migrations...")
			m, err := migrate.New(
				cfg.DBUrl,
				"file://migrations",
			)
			if err != nil {
				return err
			}
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				return err
			}
			log.Infof("Migrations complete")
			return nil
		},
	}
}
