package config

import "fmt"

const host string = "localhost"
const port int = 1323

type Config struct {
	Host string
	Port int
}

func DefaultConfig() *Config {
	return &Config{Host: host,Port: port}
}

func (cfg *Config) Addr() string {
	if cfg.Host == "" || cfg.Host == "localhost" {
		return fmt.Sprintf(":%d", cfg.Port)
	}
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}

