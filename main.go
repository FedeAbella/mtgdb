package main

import (
	"log"
	"time"
)

func main() {

	start := time.Now()

	setData, err := ReadSetList()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read set list. Got %d sets.", len(setData.Data))

	cardData, err := ReadAtomicCards()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read all atomic cards. Got %d cards.", len(cardData.Data))

	setCardData, err := ReadSetCards()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read all set cards. Got %d cards versions.", len(setCardData.Data))

	end := time.Now()
	log.Printf("Took %.2f seconds.", end.Sub(start).Seconds())
}
