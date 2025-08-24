package files

import (
	"bytes"
	"encoding/json"
	"log"
	"maps"
	"os/exec"
	"slices"

	"github.com/google/uuid"

	"FedeAbella/mtgdb/internal/datamodel"
)

func RunJQCmd(filepath string, jqFilter string) ([]byte, error) {
	jqCmd := exec.Command("jq", jqFilter, filepath)
	jqOutBuf := bytes.Buffer{}
	jqErrBuf := bytes.Buffer{}
	jqCmd.Stdout = &jqOutBuf
	err := jqCmd.Run()

	if err != nil {
		log.Println(jqErrBuf.String())
		log.Println(err)
		return []byte{}, err
	}

	return jqOutBuf.Bytes(), nil
}

func ReadSetList() (datamodel.AllSets, error) {
	jqBytes, err := RunJQCmd(allSets.Path, allSets.JQFilter)
	if err != nil {
		log.Println(err)
		return datamodel.AllSets{}, err
	}

	var allSets datamodel.AllSets
	err = json.Unmarshal(jqBytes, &allSets)
	if err != nil {
		log.Println(err)
		return datamodel.AllSets{}, err
	}

	return allSets, nil
}

func ReadAtomicCards() (datamodel.AllAtomicCards, error) {
	jqBytes, err := RunJQCmd(allAtomicCards.Path, allAtomicCards.JQFilter)
	if err != nil {
		log.Println(err)
		return datamodel.AllAtomicCards{}, err
	}

	var allCards datamodel.AllAtomicCards
	err = json.Unmarshal(jqBytes, &allCards)
	if err != nil {
		log.Println(err)
		return datamodel.AllAtomicCards{}, err
	}

	filteredCards := make(map[uuid.UUID]datamodel.AtomicCard)

	for _, card := range allCards.Data {
		_, seen := filteredCards[card.Identifiers.ScryfallOracleId]
		if seen && card.Layout == "reversible_card" {
			continue
		}

		filteredCards[card.Identifiers.ScryfallOracleId] = card
	}

	uniqueCards := datamodel.AllAtomicCards{}
	uniqueCards.Meta = allCards.Meta
	uniqueCards.Data = slices.Collect(maps.Values(filteredCards))

	return uniqueCards, nil
}

func ReadSetCards() (datamodel.AllSetCards, error) {
	jqBytes, err := RunJQCmd(allSetCards.Path, allSetCards.JQFilter)
	if err != nil {
		log.Println(err)
		return datamodel.AllSetCards{}, err
	}

	var allCards datamodel.AllSetCards
	err = json.Unmarshal(jqBytes, &allCards)
	if err != nil {
		log.Println(err)
		return datamodel.AllSetCards{}, err
	}

	return allCards, nil
}
