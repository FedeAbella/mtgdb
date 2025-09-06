package db

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"FedeAbella/mtgdb/internal/source"
	"FedeAbella/mtgdb/internal/sqlc"
)

func mapCardsToInsertAndUpdate(
	fileCardMap map[uuid.UUID]source.CardPrinting,
	dbCards []sqlc.Card,
	now time.Time,
) ([]sqlc.InsertCardsParams, []sqlc.UpdateCardParams) {
	dbCardMap := map[string]sqlc.Card{}
	for _, dbCard := range dbCards {
		dbCardMap[dbCard.ScryfallID.String()] = dbCard
	}

	cardsToInsert := make([]sqlc.InsertCardsParams, 0)
	cardsToUpdate := make([]sqlc.UpdateCardParams, 0)

	for fileCardId, fileCard := range fileCardMap {
		dbCard, inDb := dbCardMap[fileCardId.String()]
		if !inDb {
			cardsToInsert = append(cardsToInsert, fileCard.ToDbInsertCard(now))
			continue
		}

		if !fileCard.Equals(&dbCard) {
			cardsToUpdate = append(cardsToUpdate, fileCard.ToDbUpdateCard(now))
		}
	}

	return cardsToInsert, cardsToUpdate
}

func (db *DbConf) insertCards(cardsToInsert []sqlc.InsertCardsParams) error {
	if len(cardsToInsert) == 0 {
		return nil
	}

	insertStart := time.Now()
	if _, err := db.Queries.InsertCards(context.Background(), cardsToInsert); err != nil {
		log.Println(err)
		return err
	}

	log.Printf(
		"inserted %d cards into db in %.3f seconds",
		len(cardsToInsert),
		time.Since(insertStart).Seconds(),
	)

	return nil
}

func (db *DbConf) updateCards(cardsToUpdate []sqlc.UpdateCardParams) error {
	if len(cardsToUpdate) == 0 {
		return nil
	}

	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback(context.Background())

	txq := db.Queries.WithTx(tx)
	log.Println("Starting db card update transaction")

	updateStart := time.Now()
	for _, card := range cardsToUpdate {
		if err := txq.UpdateCard(context.Background(), card); err != nil {
			log.Println(err)
			return err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		log.Println(err)
		return err
	}

	log.Printf(
		"updated %d cards into db in %.3f seconds",
		len(cardsToUpdate),
		time.Since(updateStart).Seconds(),
	)

	return nil
}

func (db *DbConf) upsertCards(fileCardMap map[uuid.UUID]source.CardPrinting) error {

	dbCards, err := db.Queries.GetAllCards(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	cardsToInsert, cardsToUpdate := mapCardsToInsertAndUpdate(fileCardMap, dbCards, time.Now())

	log.Printf("%d cards to be inserted in db", len(cardsToInsert))
	log.Printf("%d cards to be updated in db", len(cardsToUpdate))

	if err = db.insertCards(cardsToInsert); err != nil {
		log.Println(err)
		return err
	}

	if err = db.updateCards(cardsToUpdate); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
