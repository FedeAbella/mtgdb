package source

import "github.com/google/uuid"

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
