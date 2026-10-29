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
	Load() (*pgxpool.Pool, string)
}

type ConfigImpl struct {
}

func NewConfig() Config {
	return &ConfigImpl{}
}

func (c *ConfigImpl) Load() (*pgxpool.Pool, string) {
	loadErr := godotenv.Load(".env")

	appPort := os.Getenv("APP_PORT")

	if loadErr != nil {
		return &pgxpool.Pool{}, appPort
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

	return pool, appPort
}
