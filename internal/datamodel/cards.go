package datamodel

import (
	"github.com/google/uuid"

	"FedeAbella/mtgdb/internal/sqlc"
)

type AtomicCard struct {
	Name        string      `json:"name"`
	Layout      string      `json:"layout"`
	Identifiers Identifiers `json:"identifiers"`
}

type SetCard struct {
	UUID        uuid.UUID   `json:"uuid"`
	Identifiers Identifiers `json:"identifiers"`
	Name        string      `json:"name"`
	Set         string      `json:"setCode"`
	NumberInSet string      `json:"number"`
}

func (c *AtomicCard) Equals(dbCard *sqlc.Card) bool {
	return c.Name == dbCard.Name && c.Identifiers.ScryfallOracleId == dbCard.ScryfallOracleID.Bytes
}
