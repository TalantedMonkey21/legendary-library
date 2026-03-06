package config

import "os"

type Config struct {
	Port string
}

func Load() *Config {
	cfg := &Config{
		Port: getEnv("Port", ":8080"),
	}
	return cfg
}

func getEnv(key, defValue string) string {
	if key == "" {
		return defValue
	}
	value := os.Getenv(key)
	if value == "" {
		return defValue
	}
	return value
}