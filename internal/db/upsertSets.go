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

func buildSetsToInsertAndUpdate(
	fileSetMap map[uuid.UUID]source.Set,
	dbSets []sqlc.Set,
) ([]sqlc.InsertSetsParams, []sqlc.UpdateSetParams) {
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

	return setsToInsert, setsToUpdate
}

func (db *DbConf) insertSets(setsToInsert []sqlc.InsertSetsParams) error {
	if len(setsToInsert) == 0 {
		return nil
	}

	insertStart := time.Now()
	if _, err := db.Queries.InsertSets(context.Background(), setsToInsert); err != nil {
		log.Println(err)
		return err

	}

	log.Printf(
		"inserted %d sets into db in %.3f seconds",
		len(setsToInsert),
		time.Since(insertStart).Seconds(),
	)

	return nil
}

func (db *DbConf) updateSets(setsToUpdate []sqlc.UpdateSetParams) error {
	if len(setsToUpdate) == 0 {
		return nil
	}

	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback(context.Background())

	txq := db.Queries.WithTx(tx)
	log.Println("Starting db set update transaction")

	updateStart := time.Now()
	for _, set := range setsToUpdate {
		if err := txq.UpdateSet(context.Background(), set); err != nil {
			log.Println(err)
			return err
		}
	}

	if err = tx.Commit(context.Background()); err != nil {
		log.Println(err)
		return err
	}

	log.Printf(
		"updated %d sets into db in %.3f seconds",
		len(setsToUpdate),
		time.Since(updateStart).Seconds(),
	)

	return nil
}

func (db *DbConf) upsertSets(fileSetMap map[uuid.UUID]source.Set) error {
	dbSets, err := db.Queries.GetAllSets(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	setsToInsert, setsToUpdate := buildSetsToInsertAndUpdate(fileSetMap, dbSets)

	log.Printf("%d sets to be inserted in db", len(setsToInsert))
	log.Printf("%d sets to be updated in db", len(setsToUpdate))

	if err = db.insertSets(setsToInsert); err != nil {
		log.Println(err)
		return err
	}

	if err = db.updateSets(setsToUpdate); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
