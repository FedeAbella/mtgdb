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

func (conn *DbConf) updateSets(fileSetMap map[uuid.UUID]source.Set) []error {
	dbSets, err := conn.Queries.GetAllSets(context.Background())
	if err != nil {
		log.Println(err)
		return []error{err}
	}

	dbSetMap := map[string]sqlc.Set{}
	for _, dbSet := range dbSets {
		dbSetMap[dbSet.ScryfallID.String()] = dbSet
	}

	setsToInsert := make([]sqlc.InsertSetsParams, 0)
	setsToUpdate := make([]sqlc.UpdateSetParams, 0)

	for fileSetID, fileSet := range fileSetMap {
		dbSet, inDb := dbSetMap[fileSetID.String()]
		if !inDb {
			setsToInsert = append(setsToInsert, sqlc.InsertSetsParams{
				ScryfallID: pgtype.UUID{
					Bytes: fileSet.ScryfallId,
					Valid: true,
				},
				Code: fileSet.Code,
				Name: fileSet.Name,
				CreatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
			continue
		}

		if !fileSet.Equals(&dbSet) {
			setsToUpdate = append(setsToUpdate, sqlc.UpdateSetParams{
				ScryfallID: dbSet.ScryfallID,
				Code:       fileSet.Code,
				Name:       fileSet.Name,
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
		}
	}

	log.Printf("%d sets to be inserted in db", len(setsToInsert))
	log.Printf("%d sets to be updated in db", len(setsToUpdate))

	if len(setsToInsert) > 0 {
		insertStart := time.Now()
		num, err := conn.Queries.InsertSets(context.Background(), setsToInsert)
		if err != nil {
			log.Println(err)
			return []error{err}
		}

		log.Printf(
			"inserted %d sets into db in %.3f seconds",
			num,
			time.Since(insertStart).Seconds(),
		)
	}

	updateErrors := make([]error, 0)
	if len(setsToUpdate) > 0 {
		updatedSets := 0
		updateStart := time.Now()
		for _, set := range setsToUpdate {
			err := conn.Queries.UpdateSet(context.Background(), set)
			if err != nil {
				updateErrors = append(updateErrors, err)
			} else {
				updatedSets++
			}
		}

		log.Printf(
			"updated %d sets into db, with %d errors, in %.3f seconds",
			updatedSets,
			len(updateErrors),
			time.Since(updateStart).Seconds(),
		)
	}

	return updateErrors
}

func (conn *DbConf) updateCards(fileCardMap map[uuid.UUID]source.CardPrinting) []error {
	dbCards, err := conn.Queries.GetAllCards(context.Background())
	if err != nil {
		log.Println(err)
		return []error{err}
	}

	dbCardMap := map[string]sqlc.Card{}
	for _, dbCard := range dbCards {
		dbCardMap[dbCard.ScryfallID.String()] = dbCard
	}

	cardsToInsert := make([]sqlc.InsertCardsParams, 0)
	cardsToUpdate := make([]sqlc.UpdateCardParams, 0)

	for fileCardId, fileCard := range fileCardMap {
		dbCard, inDb := dbCardMap[fileCardId.String()]
		if !inDb {
			cardsToInsert = append(cardsToInsert, sqlc.InsertCardsParams{
				ScryfallID: pgtype.UUID{
					Bytes: fileCard.ScryfallId,
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: fileCard.SetScryfallId,
					Valid: true,
				},
				Name:            fileCard.Name,
				CollectorNumber: fileCard.CollectorNumber,
				ColorIdentity: pgtype.Text{
					String: fileCard.ColorIdentity,
					Valid:  fileCard.ColorIdentity != "",
				},
				Colors: pgtype.Text{
					String: fileCard.Colors,
					Valid:  fileCard.Colors != "",
				},
				LanguageCode: fileCard.Language,
				SpanishName: pgtype.Text{
					String: fileCard.NameSPA,
					Valid:  fileCard.NameSPA != "",
				},
				Rarity: pgtype.Text{
					String: fileCard.Rarity,
					Valid:  fileCard.Rarity != "",
				},
				TypeLine: pgtype.Text{
					String: fileCard.TypeLine,
					Valid:  fileCard.TypeLine != "",
				},
				ScryfallApiUri: fileCard.ScryfallAPIURI,
				ScryfallWebUri: fileCard.ScryfallWebURI,
				ScryfallOracleID: pgtype.UUID{
					Bytes: fileCard.ScryfallOracleId,
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
			continue
		}

		if !fileCard.Equals(&dbCard) {
			cardsToUpdate = append(cardsToUpdate, sqlc.UpdateCardParams{
				ScryfallID: pgtype.UUID{
					Bytes: fileCard.ScryfallId,
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: fileCard.SetScryfallId,
					Valid: true,
				},
				Name:            fileCard.Name,
				CollectorNumber: fileCard.CollectorNumber,
				ColorIdentity: pgtype.Text{
					String: fileCard.ColorIdentity,
					Valid:  fileCard.ColorIdentity != "",
				},
				Colors: pgtype.Text{
					String: fileCard.Colors,
					Valid:  fileCard.Colors != "",
				},
				LanguageCode: fileCard.Language,
				SpanishName: pgtype.Text{
					String: fileCard.NameSPA,
					Valid:  fileCard.NameSPA != "",
				},
				Rarity: pgtype.Text{
					String: fileCard.Rarity,
					Valid:  fileCard.Rarity != "",
				},
				TypeLine: pgtype.Text{
					String: fileCard.TypeLine,
					Valid:  fileCard.TypeLine != "",
				},
				ScryfallApiUri: fileCard.ScryfallAPIURI,
				ScryfallWebUri: fileCard.ScryfallWebURI,
				ScryfallOracleID: pgtype.UUID{
					Bytes: fileCard.ScryfallOracleId,
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
		}
	}

	log.Printf("%d cards to be inserted in db", len(cardsToInsert))
	log.Printf("%d cards to be updated in db", len(cardsToUpdate))

	if len(cardsToInsert) > 0 {
		insertStart := time.Now()
		num, err := conn.Queries.InsertCards(context.Background(), cardsToInsert)
		if err != nil {
			log.Println(err)
			return []error{err}
		}

		log.Printf(
			"inserted %d cards into db in %.3f seconds",
			num,
			time.Since(insertStart).Seconds(),
		)
	}

	updateErrors := make([]error, 0)
	if len(cardsToUpdate) > 0 {
		updatedCards := 0
		updateStart := time.Now()
		for _, card := range cardsToUpdate {
			err := conn.Queries.UpdateCard(context.Background(), card)
			if err != nil {
				updateErrors = append(updateErrors, err)
			} else {
				updatedCards++
			}
		}

		log.Printf(
			"updated %d cards into db, with %d errors, in %.3f seconds",
			updatedCards,
			len(updateErrors),
			time.Since(updateStart).Seconds(),
		)
	}

	return updateErrors
}

func (conn *DbConf) UpdateSetsAndCards() []error {
	setMap, cardMap, err := source.GetScryfallData()
	if err != nil {
		log.Println(err)
		return []error{err}
	}

	setErrors := conn.updateSets(setMap)
	if len(setErrors) > 0 {
		log.Printf("%v", setErrors)
		return setErrors
	}

	cardErrors := conn.updateCards(cardMap)
	if len(cardErrors) > 0 {
		log.Printf("%v", cardErrors)
		return cardErrors
	}

	return []error{}
}
