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
	// start := time.Now()
	//
	// setData, err := ReadSetList()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// setReadTime := time.Now()
	// log.Printf(
	// 	"Read set list. Got %d sets in %.2f seconds.",
	// 	len(setData.Data),
	// 	setReadTime.Sub(start).Seconds(),
	// )
	//
	// cardData, err := ReadAtomicCards()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// atomicReadTime := time.Now()
	// log.Printf(
	// 	"Read all atomic cards. Got %d cards in %.2f seconds.",
	// 	len(cardData.Data),
	// 	atomicReadTime.Sub(setReadTime).Seconds(),
	// )
	//
	// setCardData, err := ReadSetCards()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// setCardReadTime := time.Now()
	// log.Printf(
	// 	"Read all set cards. Got %d cards versions in %.2f seconds.",
	// 	len(setCardData.Data),
	// 	setCardReadTime.Sub(atomicReadTime).Seconds(),
	// )
	//
	// end := time.Now()
	// log.Printf("Took %.2f seconds to read sets, cards, and set-cards", end.Sub(start).Seconds())
}
