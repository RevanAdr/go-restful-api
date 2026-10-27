package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	Port     int
	Db       string
	Username string
	Password string
}

func Load() *pgxpool.Pool {
	enverr := godotenv.Load(".env")

	if enverr != nil {
		log.Fatal("Could not load .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	env := Env{
		Host:     "localhost",
		Port:     port,
		Db:       os.Getenv("POSTGRES_DB"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", env.Username, env.Password, env.Host, env.Port, env.Db)

	pool, poolerr := pgxpool.New(context.Background(), db_url)

	if poolerr != nil {
		log.Fatalf("Could not do configuration : %s", poolerr)
		os.Exit(1)
	}

	pingerr := pool.Ping(context.Background())

	if pingerr != nil {
		log.Fatalf("There's a ping error : %s", pingerr)
		os.Exit(1)
	}

	return pool
}
