package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDatabase(dsn string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		log.Fatalln("Unable to parse database URL", err)
	}

	DB, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Fatalln("Unable to connect with the database", err)
	}

	err = DB.Ping(ctx)
	if err != nil {
		log.Fatalln("Unable to connect with the database", err)
	}

	return DB
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
