package source

import (
	"encoding/json"
	"log"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"
)

const SCRYFALL_ALL_CARDS_PATH = "./src/all-cards.json"

func readScryfallCardsFile() ([]ScryfallCard, error) {
	file, err := os.Open(SCRYFALL_ALL_CARDS_PATH)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	cards := make([]ScryfallCard, 0)
	decoder := json.NewDecoder(file)
	_, err = decoder.Token()
	if err != nil {
		log.Println(err)
	}

	for decoder.More() {
		var card ScryfallCard
		err := decoder.Decode(&card)
		if err != nil {
			log.Println(err)
		}

		if !slices.Contains(card.Games, GamePaper) {
			continue
		}

		if card.LanguageCode != English && card.LanguageCode != Spanish {
			continue
		}

		cards = append(cards, card)
	}

	_, err = decoder.Token()
	if err != nil {
		log.Println(err)
	}

	return cards, nil
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
