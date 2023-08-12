package main

import (
	"os"
)

type Config struct {
	applicationHost   string
	applicationPort   string
	mongodbConnString string
}

func (c *Config) getConfiguration() {
	c.applicationHost = findEnv("HOST", "127.0.0.1")
	c.applicationPort = findEnv("PORT", "9999")
	c.mongodbConnString = findEnv("MONGODB_URL", "mongodb://localhost:27017")
}

func findEnv(env, def string) string {
	if os.Getenv(env) == "" {
		os.Setenv(env, def)
	}
	return os.Getenv(env)
}
