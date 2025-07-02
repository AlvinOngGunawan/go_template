package config

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type AppConfig struct {
	Port      string `env:"PORT" envDefault:"8080"`
	DBHost    string `env:"DB_HOST"`
	DBPort    string `env:"DB_PORT"`
	DBUser    string `env:"DB_USER"`
	DBPass    string `env:"DB_PASS"`
	DBName    string `env:"DB_NAME"`
	JWTSecret string `env:"JWT_SECRET"`
}
