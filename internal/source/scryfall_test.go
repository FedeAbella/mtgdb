package source

import (
	"testing"

	"github.com/google/uuid"
)

func Test_GetColors(t *testing.T) {
	cases := []struct {
		Input    ScryfallCard
		Expected string
	}{
		{
			Input: ScryfallCard{
				Colors: []Color{White},
			},
			Expected: White,
		},
		{
			Input: ScryfallCard{
				Colors: []Color{Black, Green, Red, Blue, White},
			},
			Expected: "BGRUW",
		},
		{
			Input: ScryfallCard{
				Colors: []Color{},
				Faces:  []ScryfallCardFace{},
			},
			Expected: "",
		},
		{
			Input: ScryfallCard{
				Colors: []Color{},
				Faces: []ScryfallCardFace{
					{
						Colors: []Color{White},
					},
					{
						Colors: []Color{Red},
					},
				},
			},
			Expected: "RW",
		},
		{
			Input: ScryfallCard{
				Colors: []Color{},
				Faces: []ScryfallCardFace{
					{
						Colors: []Color{Blue, White},
					},
					{
						Colors: []Color{Red, White},
					},
				},
			},
			Expected: "RUW",
		},
		{
			Input: ScryfallCard{
				Colors: []Color{},
				Faces: []ScryfallCardFace{
					{
						Colors: []Color{},
					},
					{
						Colors: []Color{Red, White},
					},
				},
			},
			Expected: "RW",
		},
		{
			Input: ScryfallCard{
				Colors: []Color{},
				Faces: []ScryfallCardFace{
					{
						Colors: []Color{Black, Green},
					},
					{
						Colors: []Color{},
					},
				},
			},
			Expected: "BG",
		},
		{
			Input: ScryfallCard{
				Colors: []Color{},
				Faces: []ScryfallCardFace{
					{
						Colors: []Color{},
					},
					{
						Colors: []Color{},
					},
				},
			},
			Expected: "",
		},
	}

	for _, tCase := range cases {
		if tCase.Input.getColors() != tCase.Expected {
			t.Fatalf(
				"Expected colors %s from card %v but got %s",
				tCase.Expected,
				tCase.Input,
				tCase.Input.getColors(),
			)
		}
	}
}

func Test_GetSpanishName(t *testing.T) {
	cases := []struct {
		Input    ScryfallCard
		Expected string
	}{
		{
			Input: ScryfallCard{
				LanguageCode: "",
			},
			Expected: "",
		},
		{
			Input: ScryfallCard{
				LanguageCode: Spanish,
				PrintedName:  "Bosque",
			},
			Expected: "Bosque",
		},
		{
			Input: ScryfallCard{
				LanguageCode: Spanish,
				PrintedName:  "",
				Faces:        []ScryfallCardFace{},
			},
			Expected: "",
		},
		{
			Input: ScryfallCard{
				LanguageCode: Spanish,
				PrintedName:  "",
				Faces: []ScryfallCardFace{
					{
						PrintedName: "",
					},
					{
						PrintedName: "",
					},
				},
			},
			Expected: "",
		},
		{
			Input: ScryfallCard{
				LanguageCode: Spanish,
				Faces: []ScryfallCardFace{
					{
						PrintedName: "Vida",
					},
					{
						PrintedName: "Muerte",
					},
				},
			},
			Expected: "Vida // Muerte",
		},
	}

	for _, tCase := range cases {
		if tCase.Input.getSpanishName() != tCase.Expected {
			t.Fatalf(
				"Expected spanish name %s from card %v but got %s",
				tCase.Expected,
				tCase.Input,
				tCase.Input.getSpanishName(),
			)
		}
	}
}

func Test_Unpack(t *testing.T) {

	cases := []struct {
		Input        ScryfallCard
		ExpectedSet  Set
		ExpectedCard CardPrinting
	}{
		{
			Input: ScryfallCard{
				CMC:              5.0,
				CollectorNumber:  "94",
				ColorIdentity:    []Color{Black, Green, Red, Blue, White},
				Colors:           []Color{Black, Green, Red, Blue, White},
				Faces:            []ScryfallCardFace{},
				LanguageCode:     English,
				Name:             "Cromat",
				PrintedName:      "",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
				ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
				ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
				ScryfallSetId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
				SetCode:          "apc",
				SetName:          "Apocalypse",
				TypeLine:         "Legendary Creature — Illusion",
			},
			ExpectedSet: Set{
				Code:       "apc",
				Name:       "Apocalypse",
				ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			},
			ExpectedCard: CardPrinting{
				CollectorNumber:  "94",
				ColorIdentity:    "BGRUW",
				Colors:           "BGRUW",
				Language:         English,
				Name:             "Cromat",
				NameSPA:          "",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/7d9e0a23-d2a8-40a6-9076-ed6fb539141b",
				ScryfallId:       uuid.MustParse("7d9e0a23-d2a8-40a6-9076-ed6fb539141b"),
				ScryfallOracleId: uuid.MustParse("376601b6-fe51-4e2d-8ec6-98f965d649a3"),
				ScryfallWebURI:   "https://scryfall.com/card/apc/94/cromat?utm_source=api",
				SetScryfallId:    uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				TypeLine:         "Legendary Creature — Illusion",
			},
		},
		{
			Input: ScryfallCard{
				CMC:              3.0,
				CollectorNumber:  "244",
				ColorIdentity:    []Color{Black, Green, Red, Blue, White},
				Colors:           []Color{},
				Faces:            []ScryfallCardFace{},
				LanguageCode:     English,
				Name:             "Commander's Sphere",
				PrintedName:      "",
				Rarity:           Common,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
				ScryfallId:       uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
				ScryfallOracleId: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
				ScryfallSetId:    uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
				ScryfallWebURI:   "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
				SetCode:          "dsc",
				SetName:          "Duskmourn: House of Horror Commander",
				TypeLine:         "Artifact",
			},
			ExpectedSet: Set{
				Code:       "dsc",
				Name:       "Duskmourn: House of Horror Commander",
				ScryfallId: uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
			},
			ExpectedCard: CardPrinting{
				CollectorNumber:  "244",
				ColorIdentity:    "BGRUW",
				Colors:           "",
				Language:         English,
				Name:             "Commander's Sphere",
				NameSPA:          "",
				Rarity:           Common,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/a4ce6b63-0b38-4582-94d5-c733af087038",
				ScryfallId:       uuid.MustParse("a4ce6b63-0b38-4582-94d5-c733af087038"),
				ScryfallOracleId: uuid.MustParse("0b67c4e2-f88b-4e01-85a1-9d5f5b8db13b"),
				ScryfallWebURI:   "https://scryfall.com/card/dsc/244/commanders-sphere?utm_source=api",
				SetScryfallId:    uuid.MustParse("4c822528-83c3-42c7-8708-dd1d37166819"),
				TypeLine:         "Artifact",
			},
		},
		{
			Input: ScryfallCard{
				CMC:              10.0,
				CollectorNumber:  "5",
				ColorIdentity:    []Color{},
				Colors:           []Color{},
				Faces:            []ScryfallCardFace{},
				LanguageCode:     English,
				Name:             "Ulamog, the Ceaseless Hunger",
				PrintedName:      "",
				Rarity:           Mythic,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/c74ae706-b3b3-4097-a387-6f6c38a9b603",
				ScryfallId:       uuid.MustParse("c74ae706-b3b3-4097-a387-6f6c38a9b603"),
				ScryfallOracleId: uuid.MustParse("0bfa4512-e35a-4c93-b324-80ec659f5a97"),
				ScryfallSetId:    uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
				ScryfallWebURI:   "https://scryfall.com/card/cmm/5/ulamog-the-ceaseless-hunger?utm_source=api",
				SetCode:          "cmm",
				SetName:          "Commander Masters",
				TypeLine:         "Legendary Creature — Eldrazi",
			},
			ExpectedSet: Set{
				Code:       "cmm",
				Name:       "Commander Masters",
				ScryfallId: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
			},
			ExpectedCard: CardPrinting{
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
		},
		{
			Input: ScryfallCard{
				CMC:              6.0,
				CollectorNumber:  "335",
				ColorIdentity:    []Color{Red, Blue},
				Colors:           []Color{Red, Blue},
				Faces:            []ScryfallCardFace{},
				LanguageCode:     Spanish,
				Name:             "The Locust God",
				PrintedName:      "El Dios Langosta",
				Rarity:           Mythic,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
				ScryfallId:       uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
				ScryfallOracleId: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
				ScryfallSetId:    uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
				ScryfallWebURI:   "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
				SetCode:          "moc",
				SetName:          "March of the Machine Commander",
				TypeLine:         "Legendary Creature — God",
			},
			ExpectedSet: Set{
				Code:       "moc",
				Name:       "March of the Machine Commander",
				ScryfallId: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
			},
			ExpectedCard: CardPrinting{
				CollectorNumber:  "335",
				ColorIdentity:    "RU",
				Colors:           "RU",
				Language:         Spanish,
				Name:             "The Locust God",
				NameSPA:          "El Dios Langosta",
				Rarity:           Mythic,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/bb270c8a-91e0-4264-b036-0fcdd08fc53a",
				ScryfallId:       uuid.MustParse("bb270c8a-91e0-4264-b036-0fcdd08fc53a"),
				ScryfallOracleId: uuid.MustParse("e025a714-02da-4b0c-8021-cf3e8dc9b19e"),
				ScryfallWebURI:   "https://scryfall.com/card/moc/335/es/el-dios-langosta-(the-locust-god)?utm_source=api",
				SetScryfallId:    uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
				TypeLine:         "Legendary Creature — God",
			},
		},
		{
			Input: ScryfallCard{
				CMC:             3.0,
				CollectorNumber: "189",
				ColorIdentity:   []Color{Green},
				Colors:          []Color{},
				Faces: []ScryfallCardFace{
					{
						Colors:      []Color{Green},
						PrintedName: "Nissa, vidente del Bosque Extenso",
					},
					{
						Colors:      []Color{Green},
						PrintedName: "Nissa, animista sabia",
					},
				},
				LanguageCode:     Spanish,
				Name:             "Nissa, Vastwood Seer // Nissa, Sage Animist",
				PrintedName:      "",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/4dea9d98-1bd4-4362-9768-67827dc28d3b",
				ScryfallId:       uuid.MustParse("4dea9d98-1bd4-4362-9768-67827dc28d3b"),
				ScryfallOracleId: uuid.MustParse("35754a21-9fba-4370-a254-292918a777ba"),
				ScryfallSetId:    uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
				ScryfallWebURI:   "https://scryfall.com/card/ori/189/es/nissa-vidente-del-bosque-extenso-nissa-animista-sabia-(nissa-vastwood-seer-nissa-sage-animist)?utm_source=api",
				SetCode:          "ori",
				SetName:          "Magic Origins",
				TypeLine:         "Legendary Creature — Elf Scout // Legendary Planeswalker — Nissa",
			},
			ExpectedSet: Set{
				Code:       "ori",
				Name:       "Magic Origins",
				ScryfallId: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
			},
			ExpectedCard: CardPrinting{
				CollectorNumber:  "189",
				ColorIdentity:    "G",
				Colors:           "G",
				Language:         Spanish,
				Name:             "Nissa, Vastwood Seer // Nissa, Sage Animist",
				NameSPA:          "Nissa, vidente del Bosque Extenso // Nissa, animista sabia",
				Rarity:           Rare,
				ScryfallAPIURI:   "https://api.scryfall.com/cards/4dea9d98-1bd4-4362-9768-67827dc28d3b",
				ScryfallId:       uuid.MustParse("4dea9d98-1bd4-4362-9768-67827dc28d3b"),
				ScryfallOracleId: uuid.MustParse("35754a21-9fba-4370-a254-292918a777ba"),
				ScryfallWebURI:   "https://scryfall.com/card/ori/189/es/nissa-vidente-del-bosque-extenso-nissa-animista-sabia-(nissa-vastwood-seer-nissa-sage-animist)?utm_source=api",
				SetScryfallId:    uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
				TypeLine:         "Legendary Creature — Elf Scout // Legendary Planeswalker — Nissa",
			},
		},
	}

	for _, tCase := range cases {
		set, printing := tCase.Input.unpack()
		if set != tCase.ExpectedSet {
			t.Fatalf(
				"Expected set %v from scryfall card %v but got %v instead",
				tCase.ExpectedSet,
				tCase.Input,
				set,
			)
		}
		if printing != tCase.ExpectedCard {
			t.Fatalf(
				"Expected card %v from scryfall card %v but got %v instead",
				tCase.ExpectedCard,
				tCase.Input,
				printing,
			)
		}
	}
}
