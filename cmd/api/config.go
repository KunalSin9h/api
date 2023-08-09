package main

import (
	"os"
)

type Config struct {
	host string
	port string
}

func (c *Config) getConfiguration() error {

	if os.Getenv("host") == "" {
		os.Setenv("host", "127.0.0.1")
	}
	if os.Getenv("post") == "" {
		os.Setenv("port", "9999")
	}

	c.host = os.Getenv("host")
	c.port = os.Getenv("port")

	return nil
}
