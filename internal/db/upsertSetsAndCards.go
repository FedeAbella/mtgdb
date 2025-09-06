package db

import (
	"log"

	"FedeAbella/mtgdb/internal/source"
)

func (db *DbConf) UpsertSetsAndCards() error {
	setMap, cardMap, err := source.GetScryfallData()
	if err != nil {
		log.Println(err)
		return err
	}

	if err = db.upsertSets(setMap); err != nil {
		log.Println(err)
		return err
	}

	if err = db.upsertCards(cardMap); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
