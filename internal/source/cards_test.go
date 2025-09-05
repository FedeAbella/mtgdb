package source

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/sqlc"
)

func Test_EqualsSqlcCard(t *testing.T) {
	cases := []struct {
		Printing      CardPrinting
		SqlcCard      sqlc.Card
		ExpectedEqual bool
	}{
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: true,
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "108",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Collector Number different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRU",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Color Identity different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: English,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Language different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last tand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Name different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última esistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Spanish name different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Common,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Rarity different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a6",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // API URI different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a6"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Scryfall ID different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad55"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Scryfall Oracle ID different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-esistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Web URI different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce1"),
					Valid: true,
				},
				TypeLine: "Sorcery",
			},
			ExpectedEqual: false, // Set ID different
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcCard: sqlc.Card{
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				Name:         "Last Stand",
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				TypeLine: "Instant",
			},
			ExpectedEqual: false, // Type line different
		},
	}

	for _, tCase := range cases {
		if res := tCase.Printing.Equals(&tCase.SqlcCard); res != tCase.ExpectedEqual {
			t.Fatalf(
				"Expected equality between %v and %v to return %v but returned %v instead",
				tCase.Printing,
				tCase.SqlcCard,
				tCase.ExpectedEqual,
				res,
			)
		}
	}
}

func Test_ToDbInsertCard(t *testing.T) {
	now := time.Date(2025, 9, 4, 21, 34, 0, 0, time.UTC)

	cases := []struct {
		Printing   CardPrinting
		SqlcParams sqlc.InsertCardsParams
	}{
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcParams: sqlc.InsertCardsParams{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Name:            "Last Stand",
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				TypeLine:       "Sorcery",
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
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
			},
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "94",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         English,
				Name:             "Cromat",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
				ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
				ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Legendary Creature — Illusion",
			},
			SqlcParams: sqlc.InsertCardsParams{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Name:            "Cromat",
				CollectorNumber: "94",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: English,
				SpanishName: pgtype.Text{
					String: "",
					Valid:  false,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				TypeLine:       "Legendary Creature — Illusion",
				ScryfallApiUri: "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
				ScryfallWebUri: "https://scryfall.com/card/apc/94/cromat?utm_source=api",
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
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
			},
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "5",
				ColorIdentity:    "",
				Colors:           "",
				Language:         English,
				Name:             "Ulamog, the Ceaseless Hunger",
				NameSPA:          "",
				Rarity:           Mythic,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
				ScryfallId:       uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
				ScryfallOracleId: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
				ScryfallWebURI:   "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
				SetScryfallId:    uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
				TypeLine:         "Legendary Creature — Eldrazi",
			},
			SqlcParams: sqlc.InsertCardsParams{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
					Valid: true,
				},
				Name:            "Ulamog, the Ceaseless Hunger",
				CollectorNumber: "5",
				ColorIdentity: pgtype.Text{
					String: "",
					Valid:  false,
				},
				Colors: pgtype.Text{
					String: "",
					Valid:  false,
				},
				LanguageCode: English,
				SpanishName: pgtype.Text{
					String: "",
					Valid:  false,
				},
				Rarity: pgtype.Text{
					String: Mythic,
					Valid:  true,
				},
				TypeLine:       "Legendary Creature — Eldrazi",
				ScryfallApiUri: "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
				ScryfallWebUri: "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
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
			},
		},
	}

	for _, tCase := range cases {
		if params := tCase.Printing.ToDbInsertCard(now); params != tCase.SqlcParams {
			t.Fatalf(
				"Expected insert params %v from card %v but got %v instead",
				tCase.SqlcParams,
				tCase.Printing,
				params,
			)
		}
	}
}

func Test_ToDbUpdateCard(t *testing.T) {
	now := time.Date(2025, 9, 4, 21, 34, 0, 0, time.UTC)

	cases := []struct {
		Printing   CardPrinting
		SqlcParams sqlc.UpdateCardParams
	}{
		{
			Printing: CardPrinting{
				CollectorNumber:  "107",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         Spanish,
				Name:             "Last Stand",
				NameSPA:          "Última Resistencia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallId:       uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
				ScryfallOracleId: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Sorcery",
			},
			SqlcParams: sqlc.UpdateCardParams{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("47fee476-25b6-40bb-afa9-d755c9a021a5"),
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Name:            "Last Stand",
				CollectorNumber: "107",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: Spanish,
				SpanishName: pgtype.Text{
					String: "Última Resistencia",
					Valid:  true,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				TypeLine:       "Sorcery",
				ScryfallApiUri: "https://api.scryfall.com/cards/47fee476-25b6-40bb-afa9-d755c9a021a5",
				ScryfallWebUri: "https://scryfall.com/card/apc/107/es/ultima-resistencia-(last-stand)?utm_source=api",
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("4d2a465e-9ebd-4002-b6cd-e0eab08bad54"),
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  now,
					Valid: true,
				},
			},
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "94",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         English,
				Name:             "Cromat",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
				ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
				ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Legendary Creature — Illusion",
			},
			SqlcParams: sqlc.UpdateCardParams{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Name:            "Cromat",
				CollectorNumber: "94",
				ColorIdentity: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				Colors: pgtype.Text{
					String: "BGRUW",
					Valid:  true,
				},
				LanguageCode: English,
				SpanishName: pgtype.Text{
					String: "",
					Valid:  false,
				},
				Rarity: pgtype.Text{
					String: Rare,
					Valid:  true,
				},
				TypeLine:       "Legendary Creature — Illusion",
				ScryfallApiUri: "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
				ScryfallWebUri: "https://scryfall.com/card/apc/94/cromat?utm_source=api",
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  now,
					Valid: true,
				},
			},
		},
		{
			Printing: CardPrinting{
				CollectorNumber:  "5",
				ColorIdentity:    "",
				Colors:           "",
				Language:         English,
				Name:             "Ulamog, the Ceaseless Hunger",
				NameSPA:          "",
				Rarity:           Mythic,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
				ScryfallId:       uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
				ScryfallOracleId: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
				ScryfallWebURI:   "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
				SetScryfallId:    uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
				TypeLine:         "Legendary Creature — Eldrazi",
			},
			SqlcParams: sqlc.UpdateCardParams{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
					Valid: true,
				},
				SetID: pgtype.UUID{
					Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
					Valid: true,
				},
				Name:            "Ulamog, the Ceaseless Hunger",
				CollectorNumber: "5",
				ColorIdentity: pgtype.Text{
					String: "",
					Valid:  false,
				},
				Colors: pgtype.Text{
					String: "",
					Valid:  false,
				},
				LanguageCode: English,
				SpanishName: pgtype.Text{
					String: "",
					Valid:  false,
				},
				Rarity: pgtype.Text{
					String: Mythic,
					Valid:  true,
				},
				TypeLine:       "Legendary Creature — Eldrazi",
				ScryfallApiUri: "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
				ScryfallWebUri: "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
				ScryfallOracleID: pgtype.UUID{
					Bytes: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
					Valid: true,
				},
				UpdatedAt: pgtype.Timestamp{
					Time:  now,
					Valid: true,
				},
			},
		},
	}

	for _, tCase := range cases {
		if params := tCase.Printing.ToDbUpdateCard(now); params != tCase.SqlcParams {
			t.Fatalf(
				"Expected update params %v from card %v but got %v instead",
				tCase.SqlcParams,
				tCase.Printing,
				params,
			)
		}
	}
}
