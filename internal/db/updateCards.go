package db

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/source"
	"FedeAbella/mtgdb/internal/sqlc"
)

func (conn *DbConf) UpdateCards() []error {
	log.Printf("Updating All Atomic Cards")

	start := time.Now()
	fileCards, err := source.ReadAtomicCards()
	if err != nil {
		log.Println(err)
		return []error{err}
	}
	log.Printf(
		"Read AllAtomicCards file, got %d cards in %.3f seconds",
		len(fileCards.Data),
		time.Since(start).Seconds(),
	)

	dbStartTime := time.Now()
	dbCards, err := conn.Queries.GetAllCards(context.Background())
	if err != nil {
		log.Println(err)
		return []error{err}
	}
	log.Printf(
		"Read DB, got %d cards in %.3f seconds",
		len(dbCards),
		time.Since(dbStartTime).Seconds(),
	)

	dbCardMap := make(map[string]sqlc.Card)
	for _, card := range dbCards {
		dbCardMap[card.ScryfallOracleID.String()] = card
	}

	insertCards := make([]sqlc.InsertCardsParams, 0)
	updateCards := make([]sqlc.UpdateCardParams, 0)
	for _, fCard := range fileCards.Data {
		dbCard, inDB := dbCardMap[fCard.Identifiers.ScryfallOracleId.String()]
		if !inDB {
			insertCards = append(insertCards, sqlc.InsertCardsParams{
				ID: pgtype.UUID{
					Bytes: uuid.New(),
					Valid: true,
				},
				Name: fCard.Name,
				ScryfallOracleID: pgtype.UUID{
					Bytes: fCard.Identifiers.ScryfallOracleId,
					Valid: true,
				},
				CreatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
		} else if !fCard.Equals(&dbCard) {
			updateCards = append(updateCards, sqlc.UpdateCardParams{
				ID:   dbCard.ID,
				Name: fCard.Name,
				ScryfallOracleID: pgtype.UUID{
					Bytes: fCard.Identifiers.ScryfallOracleId,
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
		}
	}

	log.Printf("%d cards need to be inserted", len(insertCards))
	log.Printf("%d cards need to be updated", len(updateCards))

	if len(insertCards) > 0 {
		insertStartTime := time.Now()

		num, err := conn.Queries.InsertCards(context.Background(), insertCards)
		if err != nil {
			log.Println(err)
			return []error{err}
		}

		log.Printf(
			"Inserted %d cards into db in %.3f seconds",
			num,
			time.Since(insertStartTime).Seconds(),
		)
	}

	updateErrors := make([]error, 0)
	if len(updateCards) > 0 {
		updateStartTime := time.Now()
		for _, set := range updateCards {
			err := conn.Queries.UpdateCard(context.Background(), set)
			if err != nil {
				log.Println(err)
				updateErrors = append(updateErrors, err)
			}
		}
		log.Printf(
			"Updated %d cards on db in %.3f seconds with %d errors",
			len(updateCards),
			time.Since(updateStartTime).Seconds(),
			len(updateErrors),
		)
	}

	log.Printf("Completed update cards operation in %.3f seconds", time.Since(start).Seconds())

	if len(updateErrors) > 0 {
		return updateErrors
	}
	return nil
}
