package main

import (
	"context"
	"log/slog"
	"os"
	"person-details-service/internal/app/person"
	"person-details-service/internal/infrastructure/age"
	"person-details-service/internal/infrastructure/gender"
	"person-details-service/internal/infrastructure/nationality"
	"person-details-service/internal/infrastructure/person"
	"person-details-service/internal/service/person"
	"person-details-service/pkg/postgres"
	"strconv"
	"time"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Launch application")

	pgCfg, err := loadEnvConfig(logger)
	if err != nil {
		logger.Error("Unable load config",
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Info("Connect to database")

	pgPool, err := postgres.NewPool(ctx, pgCfg, logger)
	if err != nil {
		logger.Error("Unable to connect to database",
			slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer pgPool.Close()

	logger.Info("Successfully connection to database")

	if err := pgPool.Ping(ctx); err != nil {
		logger.Error("error checking the database connection",
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Connection to database has been successfully verified")

	ageRepo, err := age_infra.NewAgeRepository(*logger, age_infra.BASE_URL, 10*time.Second)
	if err != nil {
		logger.Error("Unable to create age repository",
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	genderRepo, err := gender_infra.NewGenderRepository(*logger, gender_infra.BASE_URL, 10*time.Second)
	if err != nil {
		logger.Error("Unable to create gender repository",
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	nationalityRepo, err := nationality_infra.NewNationalityRepository(*logger, nationality_infra.BASE_URL, 10*time.Second)
	if err != nil {
		logger.Error("Unable to create nationality repository",
			slog.String("error", err.Error()))
		os.Exit(1)
	}

	personRepo := person_infra.NewPersonRepository(pgPool)

	personService := person_service.NewPersonService(ageRepo, genderRepo, nationalityRepo, personRepo)

	personHandler := person_app.NewPersonHandler(ctx, personService, logger)

	router := httprouter.New()

	personHandler.Register(router)

	appPort, exists := os.LookupEnv("APP_PORT")
	if !exists {
		appPort = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", appPort)

	logger.Info("Starting server", slog.String("address", serverAddr))

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		logger.Error("Server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Application has been successfully launched")
}

func loadEnvConfig(logger *slog.Logger) (postgres.PgConfig, error) {
	config := postgres.PgConfig{
		Username:     "postgres",
		Password:     "postgres",
		Host:         "localhost",
		Port:         "5432",
		Database:     "personsdb",
		PoolMaxConns: 4,
		MaxAttempts:  3,
		MaxDelay:     10 * time.Second,
		SSLMode:      "disable",
	}

	err := godotenv.Load("../pds.env")
	if err != nil {
		logger.Warn("Unable to load .env file, using values by default",
			slog.String("error", err.Error()))
	} else {
		logger.Info(".env file loaded successfully")
	}

	if val, exists := os.LookupEnv("DB_USER"); exists {
		config.Username = val
	}

	if val, exists := os.LookupEnv("DB_PASSWORD"); exists {
		config.Password = val
	}

	if val, exists := os.LookupEnv("DB_HOST"); exists {
		config.Host = val
	}

	if val, exists := os.LookupEnv("DB_PORT"); exists {
		config.Port = val
	}

	if val, exists := os.LookupEnv("DB_NAME"); exists {
		config.Database = val
	}

	if val, exists := os.LookupEnv("DB_POOL_MAX_CONNS"); exists {
		if intVal, err := strconv.Atoi(val); err == nil {
			config.PoolMaxConns = intVal
		} else {
			logger.Warn("Invalid DB_POOL_MAX_CONNS value, the default value is used",
				slog.String("error", err.Error()),
				slog.Int("default_value", config.PoolMaxConns))
		}
	}

	if val, exists := os.LookupEnv("DB_MAX_ATTEMPTS"); exists {
		if intVal, err := strconv.Atoi(val); err == nil {
			config.MaxAttempts = intVal
		} else {
			logger.Warn("Invalid DB_MAX_ATTEMPTS value, the default value is used",
				slog.String("error", err.Error()),
				slog.Int("default_value", config.MaxAttempts))
		}
	}

	if val, exists := os.LookupEnv("DB_MAX_DELAY"); exists {
		if duration, err := time.ParseDuration(val); err == nil {
			config.MaxDelay = duration
		} else {
			logger.Warn("Invalid DB_MAX_DELAY value, the default value is used",
				slog.String("error", err.Error()),
				slog.Duration("default_value", config.MaxDelay))
		}
	}

	if val, exists := os.LookupEnv("DB_SSL_MODE"); exists {
		config.SSLMode = val
	}

	maskedPassword := "******"
	if config.Password == "" {
		maskedPassword = "<empty>"
	}

	logger.Debug("PostgreSQL config loaded",
		slog.String("username", config.Username),
		slog.String("password", maskedPassword),
		slog.String("host", config.Host),
		slog.String("port", config.Port),
		slog.String("database", config.Database),
		slog.Int("pool_max_conns", config.PoolMaxConns),
		slog.Int("max_attempts", config.MaxAttempts),
		slog.Duration("max_delay", config.MaxDelay),
		slog.String("ssl_mode", config.SSLMode))

	return config, nil
}
