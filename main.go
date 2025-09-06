package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"

	"FedeAbella/mtgdb/internal/db"
	"FedeAbella/mtgdb/internal/sqlc"
)

func main() {

	_ = godotenv.Load()
	conn, err := pgx.Connect(context.Background(), os.Getenv("GO_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	db := db.DbConf{
		Conn:    conn,
		Queries: sqlc.New(conn),
	}

	if err = db.UpsertSetsAndCards(); err != nil {
		log.Fatal(err)
	}
}
