package main

import (
	"os"
)

type Config struct {
	host string
	port string
}

func (c *Config) getConfiguration() {
	if os.Getenv("HOST") == "" {
		os.Setenv("HOST", "127.0.0.1")
	}
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "9999")
	}

	c.host = os.Getenv("HOST")
	c.port = os.Getenv("PORT")
}
