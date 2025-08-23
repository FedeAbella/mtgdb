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
	setReadTime := time.Now()
	log.Printf(
		"Read set list. Got %d sets in %.2f seconds.",
		len(setData.Data),
		setReadTime.Sub(start).Seconds(),
	)

	cardData, err := ReadAtomicCards()
	if err != nil {
		log.Fatal(err)
	}
	atomicReadTime := time.Now()
	log.Printf(
		"Read all atomic cards. Got %d cards in %.2f seconds.",
		len(cardData.Data),
		atomicReadTime.Sub(setReadTime).Seconds(),
	)

	setCardData, err := ReadSetCards()
	if err != nil {
		log.Fatal(err)
	}
	setCardReadTime := time.Now()
	log.Printf(
		"Read all set cards. Got %d cards versions in %.2f seconds.",
		len(setCardData.Data),
		setCardReadTime.Sub(atomicReadTime).Seconds(),
	)

	end := time.Now()
	log.Printf("Took %.2f seconds to read sets, cards, and set-cards", end.Sub(start).Seconds())
}
