package source

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

func readScryfallCardsFile() ([]ScryfallCard, error) {
	jqBytes, err := RunJQCmd(allScryfallCards)
	if err != nil {
		log.Println(err)
		return make([]ScryfallCard, 0), err
	}

	var sfCards []ScryfallCard
	err = json.Unmarshal(jqBytes, &sfCards)
	if err != nil {
		log.Println(err)
		return make([]ScryfallCard, 0), err
	}

	return sfCards, nil
}

func GetScryfallData() (map[uuid.UUID]Set, map[uuid.UUID]CardPrinting, error) {
	readStart := time.Now()
	sfCards, err := readScryfallCardsFile()
	if err != nil {
		log.Println(err)
		return map[uuid.UUID]Set{}, map[uuid.UUID]CardPrinting{}, err
	}
	log.Printf(
		"Read scryfall file, found %d objects in %.3f seconds",
		len(sfCards),
		time.Since(readStart).Seconds(),
	)

	setMap := make(map[uuid.UUID]Set)
	printMap := make(map[uuid.UUID]CardPrinting)

	for _, sfCard := range sfCards {
		set, printing := sfCard.unpack()
		setMap[set.ScryfallId] = set
		printMap[printing.ScryfallId] = printing
	}

	log.Printf(
		"Unpacked Scryfall data into %d sets and %d printings",
		len(setMap),
		len(printMap),
	)

	return setMap, printMap, nil
}
