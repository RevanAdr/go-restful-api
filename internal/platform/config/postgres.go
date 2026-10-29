package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Config interface {
	Load() *pgxpool.Pool
}

type ConfigImpl struct {
}

func NewConfig() Config {
	return &ConfigImpl{}
}

func (c *ConfigImpl) Load() *pgxpool.Pool {
	loadErr := godotenv.Load(".env")
	if loadErr != nil {
		return &pgxpool.Pool{}
	}

	port := ConvertPort(os.Getenv("POSTGRES_PORT"))

	env := Env{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Db:       os.Getenv("POSTGRES_DB"),
	}

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", env.User, env.Password, env.Host, env.Port, env.Db)

	pool, err := pgxpool.New(context.Background(), db_url)

	if err != nil {
		log.Fatal("Could not make database connection!")
	}

	return pool
}
