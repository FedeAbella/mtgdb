package source

import (
	"github.com/google/uuid"

	"FedeAbella/mtgdb/internal/sqlc"
)

type Set struct {
	Code       string
	Name       string
	ScryfallId uuid.UUID
}

type CardPrinting struct {
	CollectorNumber  string
	ColorIdentity    string
	Colors           string
	Language         string
	Name             string
	NameSPA          string
	Rarity           string
	ScryfallAPIURI   string
	ScryfallId       uuid.UUID
	ScryfallOracleId uuid.UUID
	ScryfallWebURI   string
	SetScryfallId    uuid.UUID
	TypeLine         string
}

func (s *Set) Equals(dbSet *sqlc.Set) bool {
	return s.ScryfallId == dbSet.ScryfallID.Bytes && s.Code == dbSet.Code && s.Name == dbSet.Name
}

func (c *CardPrinting) Equals(dbCard *sqlc.Card) bool {
	return c.CollectorNumber == dbCard.CollectorNumber &&
		c.ColorIdentity == dbCard.ColorIdentity.String &&
		c.Colors == dbCard.Colors.String &&
		c.Language == dbCard.LanguageCode &&
		c.Name == dbCard.Name &&
		c.NameSPA == dbCard.SpanishName.String &&
		c.Rarity == dbCard.Rarity.String &&
		c.ScryfallAPIURI == dbCard.ScryfallApiUri &&
		c.ScryfallId == dbCard.ScryfallID.Bytes &&
		c.ScryfallOracleId == dbCard.ScryfallOracleID.Bytes &&
		c.ScryfallWebURI == dbCard.ScryfallWebUri &&
		c.SetScryfallId == dbCard.SetID.Bytes &&
		c.TypeLine == dbCard.TypeLine.String
}
