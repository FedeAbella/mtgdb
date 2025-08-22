package datamodel

import "github.com/google/uuid"

type Identifiers struct {
	ScryfallOracleId uuid.UUID `json:"scryfallOracleId"`
	ScryfallId       uuid.UUID `json:"scryfallId"`
	MtgjsonV4Id      uuid.UUID `json:"mtgjsonV4Id"`
}
