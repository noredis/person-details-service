package postgres

import "time"

type pgConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	Database     string
	PoolMaxConns int
	MaxAttempts  int
	MaxDelay     time.Duration
	SSLMode      string
}

func NewPgConfig(
	username string,
	password string,
	host string,
	port string,
	database string,
	poolMaxConns int,
	maxAttempts int,
	maxDelay time.Duration,
	sslMode string,
) pgConfig {
	return pgConfig{
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		Database:     database,
		PoolMaxConns: poolMaxConns,
		MaxAttempts:  maxAttempts,
		MaxDelay:     maxDelay,
		SSLMode:      sslMode,
	}
}
