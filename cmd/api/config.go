package main

import (
	"os"
)

type Config struct {
	applicationHost string
	applicationPort string
	mongodbUsername string
	mongodbPassword string
	mongodbHost     string
	mongoPort       string
}

func (c *Config) getConfiguration() {
	c.applicationHost = findEnv("HOST", "127.0.0.1")
	c.applicationPort = findEnv("PORT", "9999")
	c.mongodbUsername = findEnv("MONGODB_USER", "api")
	c.mongodbPassword = findEnv("MONGODB_PASSWORD", "api")
	c.mongodbHost = findEnv("MONGODB_HOST", "mongodb")
	c.mongoPort = findEnv("MONGODB_PORT", "27017")
}

func findEnv(env, def string) string {
	if os.Getenv(env) == "" {
		os.Setenv(env, def)
	}
	return os.Getenv(env)
}
