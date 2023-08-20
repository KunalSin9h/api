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
	meiliHost         string
	meiliMasterKey    string
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

	c.meiliHost = findEnv("MEILI_HOST", "http://localhost:7700")
	c.meiliMasterKey = findEnv("MEILI_MASTER_KEY")
}

func findEnv(env string, def ...string) string {
	if os.Getenv(env) == "" {
		if len(def) == 0 {
			slog.Error("ENV '%s' is missing and no default value is configured.", env)
			os.Exit(1)
		}
		os.Setenv(env, def[0])
	}
	return os.Getenv(env)
}
