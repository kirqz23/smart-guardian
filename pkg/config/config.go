package config

// Config holds application settings loaded from environment or file.
// Add fields as needed for DB, RPC endpoints, SMTP, Redis, etc.
type Config struct {
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	DBUrl         string `mapstructure:"DATABASE_URL"`
	RPCUrl        string `mapstructure:"RPC_URL"`
	GraphQLURL    string `mapstructure:"GRAPHQL_URL"`
	MetricsPeriod int    `mapstructure:"METRICS_PERIOD"`
}

// DefaultConfig provides sensible defaults for Config fields.
func DefaultConfig() Config {
	return Config{
		LogLevel:      "info",
		DBUrl:         "postgres://localhost:5432/smart_guardian?sslmode=disable",
		RPCUrl:        "",
		GraphQLURL:    "",
		MetricsPeriod: 60,
	}
}
