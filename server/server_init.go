package server

import (
	"Test_Go/constanta"
	middleware2 "Test_Go/middleware"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func (s *ServerAttribute) InitServer() (err error) {
	//if we need to init everything after server up so we don't need to init everytime we declare service
	//ex redis, nats / kafka, rabbitmq, socket
	s.Server = echo.New()
	s.Server.Use(middleware.Recover())
	s.Server.Use(middleware2.RequestIDMiddleware())
	s.Server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: constanta.LoggerFormat,
	}))

	_ = godotenv.Load()
	s.LoadConfig()
	s.DB, err = s.InitDB()
	if err != nil {
		return
	}

	s.initRepository()

	s.initServices()

	s.initHandler()

	s.initEndpoint()

	s.InitRoutes()

	return
}

func (s *ServerAttribute) InitDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		s.Config.DBUser,
		s.Config.DBPass,
		s.Config.DBHost,
		s.Config.DBPort,
		s.Config.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations", // path to your .sql files
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migration instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration up: %w", err)
	}

	log.Println("✅ DB migrations applied")
	return nil
}
