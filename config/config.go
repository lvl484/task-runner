package config

import "os"

type Config struct {
	address string
}

func Init() *Config {
	addr := os.Getenv("ADDRESS")
	if addr == "" {
		addr = ":8099"
	}
	return &Config{address: addr}
}

func (c *Config) Address() string {
	return c.address
}
