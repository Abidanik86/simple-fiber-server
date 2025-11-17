package config

import "os"

type Config struct {
	AppPort            string
	AppEnv             string
	CORSAllowedOrigins string
}

// helper: default if empty
func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return def
}

func Load() *Config {
	return &Config{
		AppPort:            getEnv("APP_PORT", ":3000"),
		AppEnv:             getEnv("APP_ENV", "development"),
		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
	}
}
