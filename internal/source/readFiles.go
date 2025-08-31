package source

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const SCRYFALL_ALL_CARDS_PATH = "./src/all-cards.json"

type jsonDecoder[T any] struct {
}

func (d *jsonDecoder[T]) decodeArray(f io.Reader) ([]T, error) {
	arr := make([]T, 0)
	decoder := json.NewDecoder(f)

	// Remove the opening array token
	_, err := decoder.Token()
	if err != nil {
		log.Println(err)
		return []T{}, err
	}

	// Extract each structured element of the array
	for decoder.More() {
		card := new(T)
		err := decoder.Decode(card)
		if err != nil {
			log.Println(err)
			return []T{}, err
		}

		arr = append(arr, *card)
	}

	// Remove the closing array token, ensuring the array is complete
	_, err = decoder.Token()
	if err != nil {
		log.Println(err)
		return []T{}, err
	}

	return arr, nil
}

func GetScryfallData() (map[uuid.UUID]Set, map[uuid.UUID]CardPrinting, error) {
	file, err := os.Open(SCRYFALL_ALL_CARDS_PATH)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	readStart := time.Now()
	decoder := jsonDecoder[ScryfallCard]{}
	sfCards, err := decoder.decodeArray(file)
	if err != nil {
		log.Println(err)
		return map[uuid.UUID]Set{}, map[uuid.UUID]CardPrinting{}, err
	}

	log.Printf(
		"Read scryfall file, found %d objects in %.3f seconds",
		len(sfCards),
		time.Since(readStart).Seconds(),
	)

	sets, cards := scryfallToSetsCards(sfCards)
	if err != nil {
		log.Println(err)
		return map[uuid.UUID]Set{}, map[uuid.UUID]CardPrinting{}, err
	}

	log.Printf(
		"Unpacked Scryfall data into %d sets and %d printings",
		len(sets),
		len(cards),
	)

	return sets, cards, nil
}
