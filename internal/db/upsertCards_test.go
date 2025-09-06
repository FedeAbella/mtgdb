package db

import (
	"slices"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/source"
	"FedeAbella/mtgdb/internal/sqlc"
)

func Test_MapToInsertAndUpdate(t *testing.T) {
	now := time.Date(2025, 9, 5, 21, 36, 0, 0, time.UTC)
	past := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name                string
		cardsInDb           []sqlc.Card
		cardsInFile         map[uuid.UUID]source.CardPrinting
		expectedInsertCards []sqlc.InsertCardsParams
		expectedUpdateCards []sqlc.UpdateCardParams
	}{
		{
			name:      "no cards in DB",
			cardsInDb: []sqlc.Card{},
			cardsInFile: map[uuid.UUID]source.CardPrinting{
				uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"): {
					CollectorNumber:  "94",
					ColorIdentity:    "BGRUW",
					Colors:           "BGRUW",
					Language:         source.English,
					Name:             "Cromat",
					NameSPA:          "",
					Rarity:           source.Rare,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
					ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
					ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					TypeLine:         "Legendary Creature — Illusion",
				},
				uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"): {
					CollectorNumber:  "244",
					ColorIdentity:    "BGRUW",
					Colors:           "",
					Language:         source.English,
					Name:             "Commander's Sphere",
					NameSPA:          "",
					Rarity:           source.Common,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallId:       uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
					ScryfallOracleId: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
					ScryfallWebURI:   "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetScryfallId:    uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
					TypeLine:         "Artifact",
				},
				uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"): {
					CollectorNumber:  "5",
					ColorIdentity:    "",
					Colors:           "",
					Language:         source.English,
					Name:             "Ulamog, the Ceaseless Hunger",
					NameSPA:          "",
					Rarity:           source.Mythic,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
					ScryfallId:       uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
					ScryfallOracleId: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
					ScryfallWebURI:   "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
					SetScryfallId:    uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
					TypeLine:         "Legendary Creature — Eldrazi",
				},
				uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"): {
					CollectorNumber:  "335",
					ColorIdentity:    "RU",
					Colors:           "RU",
					Language:         source.Spanish,
					Name:             "The Locust God",
					NameSPA:          "El Dios Langosta",
					Rarity:           source.Mythic,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
					ScryfallId:       uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
					ScryfallOracleId: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
					ScryfallWebURI:   "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
					SetScryfallId:    uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
					TypeLine:         "Legendary Creature — God",
				},
			},
			expectedInsertCards: []sqlc.InsertCardsParams{
				{
					CollectorNumber: "94",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					LanguageCode: source.English,
					Name:         "Cromat",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Rare,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Illusion",
					CreatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
				{
					CollectorNumber: "244",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Commander's Sphere",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Common,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
						Valid: true,
					},
					TypeLine: "Artifact",
					CreatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
				{
					CollectorNumber: "5",
					ColorIdentity: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Ulamog, the Ceaseless Hunger",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Mythic,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Eldrazi",
					CreatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
				{
					CollectorNumber: "335",
					ColorIdentity: pgtype.Text{
						String: "RU",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "RU",
						Valid:  true,
					},
					LanguageCode: source.Spanish,
					Name:         "The Locust God",
					SpanishName: pgtype.Text{
						String: "El Dios Langosta",
						Valid:  true,
					},
					Rarity: pgtype.Text{
						String: source.Mythic,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — God",
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
			expectedUpdateCards: []sqlc.UpdateCardParams{},
		},
		{
			name: "all cards in DB",
			cardsInDb: []sqlc.Card{
				{
					CollectorNumber: "94",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					LanguageCode: source.English,
					Name:         "Cromat",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Rare,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Illusion",
					CreatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
				},
				{
					CollectorNumber: "244",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Commander's Sphere",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Common,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
						Valid: true,
					},
					TypeLine: "Artifact",
					CreatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
				},
				{
					CollectorNumber: "5",
					ColorIdentity: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Ulamog, the Ceaseless Hunger",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Mythic,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Eldrazi",
					CreatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
				},
				{
					CollectorNumber: "335",
					ColorIdentity: pgtype.Text{
						String: "RU",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "RU",
						Valid:  true,
					},
					LanguageCode: source.Spanish,
					Name:         "The Locust God",
					SpanishName: pgtype.Text{
						String: "El Dios Langosta",
						Valid:  true,
					},
					Rarity: pgtype.Text{
						String: source.Mythic,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — God",
					CreatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
				},
			},
			cardsInFile: map[uuid.UUID]source.CardPrinting{
				uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"): {
					CollectorNumber:  "94",
					ColorIdentity:    "BGRUW",
					Colors:           "BGRUW",
					Language:         source.English,
					Name:             "Cromat",
					NameSPA:          "",
					Rarity:           source.Rare,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
					ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
					ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					TypeLine:         "Legendary Creature — Illusion",
				},
				uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"): {
					CollectorNumber:  "244",
					ColorIdentity:    "BGRUW",
					Colors:           "",
					Language:         source.English,
					Name:             "Commander's Sphere",
					NameSPA:          "",
					Rarity:           source.Common,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallId:       uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
					ScryfallOracleId: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
					ScryfallWebURI:   "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetScryfallId:    uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
					TypeLine:         "Artifact",
				},
				uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"): {
					CollectorNumber:  "5",
					ColorIdentity:    "",
					Colors:           "",
					Language:         source.English,
					Name:             "Ulamog, the Ceaseless Hunger",
					NameSPA:          "",
					Rarity:           source.Mythic,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
					ScryfallId:       uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
					ScryfallOracleId: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
					ScryfallWebURI:   "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
					SetScryfallId:    uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
					TypeLine:         "Legendary Creature — Eldrazi",
				},
				uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"): {
					CollectorNumber:  "335",
					ColorIdentity:    "RU",
					Colors:           "RU",
					Language:         source.Spanish,
					Name:             "The Locust God",
					NameSPA:          "El Dios Langosta",
					Rarity:           source.Mythic,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
					ScryfallId:       uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
					ScryfallOracleId: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
					ScryfallWebURI:   "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
					SetScryfallId:    uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
					TypeLine:         "Legendary Creature — God",
				},
			},
			expectedInsertCards: []sqlc.InsertCardsParams{},
			expectedUpdateCards: []sqlc.UpdateCardParams{},
		},
		{
			name: "some cards to insert, some to update",
			cardsInDb: []sqlc.Card{
				{
					CollectorNumber: "94",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					LanguageCode: source.English,
					Name:         "Cromat",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Rare,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Illusion",
					CreatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
				},
				{
					CollectorNumber: "244",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Commander's Sphere",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Common,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
						Valid: true,
					},
					TypeLine: "Artifact",
					CreatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  past,
						Valid: true,
					},
				},
			},
			cardsInFile: map[uuid.UUID]source.CardPrinting{
				uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"): {
					CollectorNumber:  "94",
					ColorIdentity:    "BGRUW",
					Colors:           "BGRUW",
					Language:         source.English,
					Name:             "Cromat -- Updated",
					NameSPA:          "",
					Rarity:           source.Rare,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
					ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
					ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					TypeLine:         "Legendary Creature — Illusion",
				},
				uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"): {
					CollectorNumber:  "244",
					ColorIdentity:    "BGRUW",
					Colors:           "",
					Language:         source.English,
					Name:             "Commander's Sphere -- Updated",
					NameSPA:          "",
					Rarity:           source.Common,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallId:       uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
					ScryfallOracleId: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
					ScryfallWebURI:   "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetScryfallId:    uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
					TypeLine:         "Artifact",
				},
				uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"): {
					CollectorNumber:  "5",
					ColorIdentity:    "",
					Colors:           "",
					Language:         source.English,
					Name:             "Ulamog, the Ceaseless Hunger",
					NameSPA:          "",
					Rarity:           source.Mythic,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
					ScryfallId:       uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
					ScryfallOracleId: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
					ScryfallWebURI:   "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
					SetScryfallId:    uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
					TypeLine:         "Legendary Creature — Eldrazi",
				},
				uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"): {
					CollectorNumber:  "335",
					ColorIdentity:    "RU",
					Colors:           "RU",
					Language:         source.Spanish,
					Name:             "The Locust God",
					NameSPA:          "El Dios Langosta",
					Rarity:           source.Mythic,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
					ScryfallId:       uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
					ScryfallOracleId: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
					ScryfallWebURI:   "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
					SetScryfallId:    uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
					TypeLine:         "Legendary Creature — God",
				},
			},
			expectedInsertCards: []sqlc.InsertCardsParams{
				{
					CollectorNumber: "5",
					ColorIdentity: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Ulamog, the Ceaseless Hunger",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Mythic,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Eldrazi",
					CreatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
				{
					CollectorNumber: "335",
					ColorIdentity: pgtype.Text{
						String: "RU",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "RU",
						Valid:  true,
					},
					LanguageCode: source.Spanish,
					Name:         "The Locust God",
					SpanishName: pgtype.Text{
						String: "El Dios Langosta",
						Valid:  true,
					},
					Rarity: pgtype.Text{
						String: source.Mythic,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — God",
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
			expectedUpdateCards: []sqlc.UpdateCardParams{
				{
					CollectorNumber: "94",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					LanguageCode: source.English,
					Name:         "Cromat -- Updated",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Rare,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/apc/94/cromat?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
						Valid: true,
					},
					TypeLine: "Legendary Creature — Illusion",
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
				{
					CollectorNumber: "244",
					ColorIdentity: pgtype.Text{
						String: "BGRUW",
						Valid:  true,
					},
					Colors: pgtype.Text{
						String: "",
						Valid:  false,
					},
					LanguageCode: source.English,
					Name:         "Commander's Sphere -- Updated",
					SpanishName: pgtype.Text{
						String: "",
						Valid:  false,
					},
					Rarity: pgtype.Text{
						String: source.Common,
						Valid:  true,
					},
					ScryfallApiUri: "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
						Valid: true,
					},
					ScryfallOracleID: pgtype.UUID{
						Bytes: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
						Valid: true,
					},
					ScryfallWebUri: "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
					SetID: pgtype.UUID{
						Bytes: uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
						Valid: true,
					},
					TypeLine: "Artifact",
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotInsert, gotUpdate := mapCardsToInsertAndUpdate(test.cardsInFile, test.cardsInDb, now)
			for _, expectedInsert := range test.expectedInsertCards {
				if !slices.Contains(gotInsert, expectedInsert) {
					t.Fatalf(
						"test %s expected card %#v to be inserted, but got %#v",
						test.name,
						expectedInsert,
						gotInsert,
					)
				}
			}

			for _, expectedUpdate := range test.expectedUpdateCards {
				if !slices.Contains(gotUpdate, expectedUpdate) {
					t.Fatalf(
						"test %s expected card %#v to be updated, but got %#v",
						test.name,
						expectedUpdate,
						gotUpdate,
					)
				}
			}

		})
	}
}
