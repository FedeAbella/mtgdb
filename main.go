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

		Queries: sqlc.New(conn),
	}

	db.UpdateSets()
	db.UpdateCards()
}
