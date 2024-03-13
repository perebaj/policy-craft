// Package main is the initial point for the service policycraft.
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	api "github.com/perebaj/policycraft/api/docs"
	"github.com/perebaj/policycraft/postgres"
)

// Config have the core configuration for the service.
type Config struct {
	// PORT is the port where the service will be listening.
	PORT string
	// LogLevel is the level of the logs. Could be INFO, DEBUG, WARN or ERROR.
	LogLevel string
	LogType  string // json(for cloud environments) or text(for local environments)
	// Postgres is the configuration for the postgres database.
	Postgres postgres.Config
}

func main() {
	// Load the configuration from the environment variables.
	cfg := Config{
		PORT:     getEnvWithDefault("PORT", "8080"),
		LogLevel: getEnvWithDefault("LOG_LEVEL", "INFO"),
		LogType:  getEnvWithDefault("LOG_TYPE", "json"),
		Postgres: postgres.Config{
			URL:             os.Getenv("POLICY_CRAFT_POSTGRES_URL"),
			MaxOpenConns:    10,
			MaxIdleConns:    5,
			ConnMaxIdleTime: 1 * time.Minute,
		},
	}

	err := setUpLog(cfg)
	if err != nil {
		slog.Error("failed to set up log", "error", err)
		os.Exit(1)
	}

	db, err := postgres.OpenDB(cfg.Postgres)
	if err != nil {
		slog.Error("failed to open database", "error", err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			slog.Error("failed to close database", "error", err)
		}
	}()

	storage := postgres.NewStorage(db)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /policies", api.SavePolicyHandler(storage))

	slog.Info("starting server", "port", cfg.PORT)

	err = http.ListenAndServe(":"+cfg.PORT, mux)
	if err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}

// setUpLog initialize the logger.
func setUpLog(cfg Config) error {
	var level slog.Level
	switch cfg.LogLevel {
	case "INFO":
		level = slog.LevelInfo
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		return fmt.Errorf("invalid log level: %s", cfg.LogLevel)
	}

	var logger *slog.Logger
	if cfg.LogType == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))
	} else if cfg.LogType == "text" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		}))
	} else {
		return fmt.Errorf("invalid log type: %s", cfg.LogType)
	}

	slog.SetDefault(logger)
	return nil
}

// getEnvWithDefault returns the value of the environment variable with the given key.
// If the environment variable is not set, it returns the default value.
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
