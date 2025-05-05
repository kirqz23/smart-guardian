package config

import (
	"github.com/spf13/viper"
)

// Load reads environment variables and returns a populated Config.
func Load() (cfg Config, err error) {
	// Initialize Viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Set defaults
	defaults := DefaultConfig()
	viper.SetDefault("LOG_LEVEL", defaults.LogLevel)
	viper.SetDefault("DATABASE_URL", defaults.DBUrl)
	viper.SetDefault("RPC_URL", defaults.RPCUrl)
	viper.SetDefault("GRAPHQL_URL", defaults.GraphQLURL)
	viper.SetDefault("METRICS_PERIOD", defaults.MetricsPeriod)

	// Read from file (optional)
	_ = viper.ReadInConfig()

	// Overwrite with environment variables
	viper.AutomaticEnv()

	// Unmarshal into struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	// Convert metrics period to duration if needed
	// For example: cfg.MetricsWindow = time.Duration(cfg.MetricsPeriod) * time.Second

	return cfg, nil
}
