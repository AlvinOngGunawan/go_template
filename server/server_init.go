package server

import (
	"Test_Go/constanta"
	middleware2 "Test_Go/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
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
	s.Server.Use(middleware.CORS())

	_ = godotenv.Load()
	s.LoadConfig()
	s.DB, err = s.InitDB()
	if err != nil {
		return
	}

	err = s.initRedis()
	if err != nil {
		return
	}

	s.initModule()

	return
}

func (s *ServerAttribute) InitDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	//if err = runMigrations(db); err != nil {
	//	return nil, err
	//}

	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
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

func (s *ServerAttribute) initRedis() (err error) {
	s.Redis = redis.NewClient(&redis.Options{
		Addr:     s.Config.RedisHost + ":" + s.Config.RedisPort,
		Password: s.Config.RedisPort,
		DB:       s.Config.RedisDB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err = s.Redis.Ping(ctx).Err(); err != nil {
		//optional to connect redis
		return nil
	}

	s.RedisLock = redislock.New(s.Redis)

	return
}
