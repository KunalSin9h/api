package main

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	applicationHost   string
	applicationPort   string
	mongodbConnString string
	dbTimeout         int
}

func (c *Config) getConfiguration() {
	c.applicationHost = findEnv("HOST", "127.0.0.1")
	c.applicationPort = findEnv("PORT", "9999")
	c.mongodbConnString = findEnv("MONGODB_URL", "mongodb://localhost:27017")
	timeout, err := strconv.Atoi(findEnv("DB_TIMEOUT", "5000"))

	if err != nil {
		slog.Error("Failed to parse DB_TIMEOUT, with error :%v", err.Error())
		os.Exit(1)
	}

	c.dbTimeout = timeout
}

func findEnv(env, def string) string {
	if os.Getenv(env) == "" {
		os.Setenv(env, def)
	}
	return os.Getenv(env)
}
