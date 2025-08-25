package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/source"
	"FedeAbella/mtgdb/internal/sqlc"
)

func (conn *DbConf) UpdateSets() []error {
	log.Printf("Updating All Sets")

	start := time.Now()
	fileSets, err := source.ReadSetList()
	if err != nil {
		log.Println(err)
		return []error{err}
	}
	log.Printf(
		"Read SetList file, got %d sets in %.3f seconds",
		len(fileSets.Data),
		time.Since(start).Seconds(),
	)

	dbStartTime := time.Now()
	dbSets, err := conn.Queries.GetAllSets(context.Background())
	if err != nil {
		log.Println(err)
		return []error{err}
	}
	log.Printf(
		"Read DB, got %d sets in %.3f seconds",
		len(dbSets),
		time.Since(dbStartTime).Seconds(),
	)

	dbSetMap := make(map[string]sqlc.Set)
	for _, set := range dbSets {
		dbSetMap[set.Code] = set
	}

	insertSets := make([]sqlc.InsertSetsParams, 0)
	updateSets := make([]sqlc.UpdateSetParams, 0)
	for _, fSet := range fileSets.Data {
		dbSet, inDB := dbSetMap[fSet.Code]
		if !inDB {
			insertSets = append(insertSets, sqlc.InsertSetsParams{
				Code: fSet.Code,
				Name: fSet.Name,
				CreatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
		} else if !fSet.Equals(&dbSet) {
			updateSets = append(updateSets, sqlc.UpdateSetParams{
				Code: fSet.Code,
				Name: fSet.Name,
				UpdatedAt: pgtype.Timestamp{
					Time:  time.Now(),
					Valid: true,
				},
			})
		}
	}

	log.Printf("%d sets need to be inserted", len(insertSets))
	log.Printf("%d sets need to be updated", len(updateSets))

	if len(insertSets) > 0 {
		insertStartTime := time.Now()

		num, err := conn.Queries.InsertSets(context.Background(), insertSets)
		if err != nil {
			log.Println(err)
			return []error{err}
		}

		log.Printf(
			"Inserted %d sets into db in %.3f seconds",
			num,
			time.Since(insertStartTime).Seconds(),
		)
	}

	updateErrors := make([]error, 0)
	if len(updateSets) > 0 {
		updateStartTime := time.Now()
		for _, set := range updateSets {
			err := conn.Queries.UpdateSet(context.Background(), set)
			if err != nil {
				log.Println(err)
				updateErrors = append(updateErrors, err)
			}
		}
		log.Printf(
			"Updated %d sets on db in %.3f seconds with %d errors",
			len(updateSets),
			time.Since(updateStartTime).Seconds(),
			len(updateErrors),
		)
	}

	log.Printf("Completed update sets operation in %.3f seconds", time.Since(start).Seconds())

	if len(updateErrors) > 0 {
		return updateErrors
	}
	return nil
}
