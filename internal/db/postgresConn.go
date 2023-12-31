package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var counts int8

func SetupDBConn() *pgxpool.Pool {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready yet")
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
