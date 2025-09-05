package source

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/sqlc"
)

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
		c.TypeLine == dbCard.TypeLine
}

func (c *CardPrinting) ToDbInsertCard(now time.Time) sqlc.InsertCardsParams {
	return sqlc.InsertCardsParams{
		ScryfallID: pgtype.UUID{
			Bytes: c.ScryfallId,
			Valid: true,
		},
		SetID: pgtype.UUID{
			Bytes: c.SetScryfallId,
			Valid: true,
		},
		Name:            c.Name,
		CollectorNumber: c.CollectorNumber,
		ColorIdentity: pgtype.Text{
			String: c.ColorIdentity,
			Valid:  c.ColorIdentity != "",
		},
		Colors: pgtype.Text{
			String: c.Colors,
			Valid:  c.Colors != "",
		},
		LanguageCode: c.Language,
		SpanishName: pgtype.Text{
			String: c.NameSPA,
			Valid:  c.NameSPA != "",
		},
		Rarity: pgtype.Text{
			String: c.Rarity,
			Valid:  c.Rarity != "",
		},
		TypeLine:       c.TypeLine,
		ScryfallApiUri: c.ScryfallAPIURI,
		ScryfallWebUri: c.ScryfallWebURI,
		ScryfallOracleID: pgtype.UUID{
			Bytes: c.ScryfallOracleId,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
	}
}

func (c *CardPrinting) ToDbUpdateCard(now time.Time) sqlc.UpdateCardParams {
	return sqlc.UpdateCardParams{
		ScryfallID: pgtype.UUID{
			Bytes: c.ScryfallId,
			Valid: true,
		},
		SetID: pgtype.UUID{
			Bytes: c.SetScryfallId,
			Valid: true,
		},
		Name:            c.Name,
		CollectorNumber: c.CollectorNumber,
		ColorIdentity: pgtype.Text{
			String: c.ColorIdentity,
			Valid:  c.ColorIdentity != "",
		},
		Colors: pgtype.Text{
			String: c.Colors,
			Valid:  c.Colors != "",
		},
		LanguageCode: c.Language,
		SpanishName: pgtype.Text{
			String: c.NameSPA,
			Valid:  c.NameSPA != "",
		},
		Rarity: pgtype.Text{
			String: c.Rarity,
			Valid:  c.Rarity != "",
		},
		TypeLine:       c.TypeLine,
		ScryfallApiUri: c.ScryfallAPIURI,
		ScryfallWebUri: c.ScryfallWebURI,
		ScryfallOracleID: pgtype.UUID{
			Bytes: c.ScryfallOracleId,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
	}
}
