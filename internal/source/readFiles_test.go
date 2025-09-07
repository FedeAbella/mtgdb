package source

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func Test_DecodeScryfallArray(t *testing.T) {
	tests := []struct {
		name        string
		Input       string
		ThrowsError bool
		Expected    []ScryfallCard
	}{
		{
			name: "good json",
			Input: `
				[
				  {
				    "id": "00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "oracle_id": "2afbaa9a-c171-4a8b-90f3-5250d8498356",
				    "name": "Shardless Agent",
				    "printed_name": "Agente sin fragmento",
				    "lang": "es",
				    "uri": "https://api.scryfall.com/cards/00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "scryfall_uri": "https://scryfall.com/card/mh2/321/es/agente-sin-fragmento-(shardless-agent)?utm_source=api",
				    "cmc": 3.0,
				    "type_line": "Artifact Creature — Human Rogue",
				    "colors": [
						"G",
						"U"
				    ],
				    "color_identity": [
						"G",
						"U"
				    ],
				    "games": [
						"paper",
						"mtgo"
				    ],
				    "set_id": "c1c7eb8c-f205-40ab-a609-767cb296544e",
				    "set": "mh2",
				    "set_name": "Modern Horizons 2",
				    "collector_number": "321",
				    "rarity": "rare"
				  },
				  {
					"id": "0000419b-0bba-4488-8f7a-6194544ce91e",
					"oracle_id": "b34bb2dc-c1af-4d77-b0b3-a0fb342a5fc6",
					"name": "Forest",
					"lang": "en",
					"uri": "https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e",
					"scryfall_uri": "https://scryfall.com/card/blb/280/forest?utm_source=api",
					"cmc": 0.0,
					"type_line": "Basic Land — Forest",
					"colors": [],
					"color_identity": [
					  "G"
					],
					"games": [
					  "paper",
					  "mtgo",
					  "arena"
					],
					"set_id": "a2f58272-bba6-439d-871e-7a46686ac018",
					"set": "blb",
					"set_name": "Bloomburrow",
					"collector_number": "280",
					"rarity": "common"
				  },
     			  {
				    "id": "006cb204-b0fb-4e83-9a85-76d86522bdc0",
				    "oracle_id": "bd43b485-e42e-4c20-adcb-799b0ce18a59",
				    "name": "Jorn, God of Winter // Kaldring, the Rimestaff",
				    "lang": "es",
				    "uri": "https://api.scryfall.com/cards/006cb204-b0fb-4e83-9a85-76d86522bdc0",
				    "scryfall_uri": "https://scryfall.com/card/khm/179/es/jorn-dios-del-invierno-k%C3%A1ldring-la-varaescarcha-(jorn-god-of-winter-kaldring-the-rimestaff)?utm_source=api",
				    "cmc": 3.0,
				    "type_line": "Legendary Snow Creature — God // Legendary Snow Artifact",
				    "color_identity": [
				      "B",
				      "G",
				      "U"
				    ],
				    "card_faces": [
				      {
				        "printed_name": "Jorn, dios del invierno",
				        "colors": [
				      	  "G"
				        ]
				      },
				      {
				        "printed_name": "Káldring, la Varaescarcha",
				        "colors": [
				      	  "B",
				      	  "U"
				        ]
				      }
				    ],
				    "games": [
				      "arena",
				      "paper",
				      "mtgo"
				    ],
				    "set_id": "43057fad-b1c1-437f-bc48-0045bce6d8c9",
				    "set": "khm",
				    "set_name": "Kaldheim",
				    "collector_number": "179",
				    "rarity": "rare"
				}
			  ]`,
			ThrowsError: false,
			Expected: []ScryfallCard{
				{
					ScryfallId:       uuid.MustParse("00017e6d-bf93-4dcf-8751-f50aba77e2d2"),
					ScryfallOracleId: uuid.MustParse("2afbaa9a-c171-4a8b-90f3-5250d8498356"),
					Name:             "Shardless Agent",
					PrintedName:      "Agente sin fragmento",
					LanguageCode:     Spanish,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/00017e6d-bf93-4dcf-8751-f50aba77e2d2",
					ScryfallWebURI:   "https://scryfall.com/card/mh2/321/es/agente-sin-fragmento-(shardless-agent)?utm_source=api",
					CMC:              3.0,
					TypeLine:         "Artifact Creature — Human Rogue",
					Colors:           []Color{Green, Blue},
					ColorIdentity:    []Color{Green, Blue},
					Games:            []string{GamePaper, "mtgo"},
					ScryfallSetId:    uuid.MustParse("c1c7eb8c-f205-40ab-a609-767cb296544e"),
					SetCode:          "mh2",
					SetName:          "Modern Horizons 2",
					CollectorNumber:  "321",
					Rarity:           Rare,
				},
				{
					ScryfallId:       uuid.MustParse("0000419b-0bba-4488-8f7a-6194544ce91e"),
					ScryfallOracleId: uuid.MustParse("b34bb2dc-c1af-4d77-b0b3-a0fb342a5fc6"),
					Name:             "Forest",
					LanguageCode:     English,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/0000419b-0bba-4488-8f7a-6194544ce91e",
					ScryfallWebURI:   "https://scryfall.com/card/blb/280/forest?utm_source=api",
					CMC:              0.0,
					TypeLine:         "Basic Land — Forest",
					Colors:           []Color{},
					ColorIdentity:    []Color{Green},
					Games:            []string{GamePaper, "mtgo", "arena"},
					ScryfallSetId:    uuid.MustParse("a2f58272-bba6-439d-871e-7a46686ac018"),
					SetCode:          "blb",
					SetName:          "Bloomburrow",
					CollectorNumber:  "280",
					Rarity:           Common,
					PrintedName:      "",
				},
				{
					ScryfallId:       uuid.MustParse("006cb204-b0fb-4e83-9a85-76d86522bdc0"),
					ScryfallOracleId: uuid.MustParse("bd43b485-e42e-4c20-adcb-799b0ce18a59"),
					Name:             "Jorn, God of Winter // Kaldring, the Rimestaff",
					LanguageCode:     Spanish,
					ScryfallAPIURI:   "https://api.scryfall.com/cards/006cb204-b0fb-4e83-9a85-76d86522bdc0",
					ScryfallWebURI:   "https://scryfall.com/card/khm/179/es/jorn-dios-del-invierno-k%C3%A1ldring-la-varaescarcha-(jorn-god-of-winter-kaldring-the-rimestaff)?utm_source=api",
					CMC:              3.0,
					TypeLine:         "Legendary Snow Creature — God // Legendary Snow Artifact",
					ColorIdentity:    []Color{Black, Green, Blue},
					Games:            []string{"arena", GamePaper, "mtgo"},
					ScryfallSetId:    uuid.MustParse("43057fad-b1c1-437f-bc48-0045bce6d8c9"),
					SetCode:          "khm",
					SetName:          "Kaldheim",
					CollectorNumber:  "179",
					Rarity:           Rare,
					Faces: []ScryfallCardFace{
						{
							PrintedName: "Jorn, dios del invierno",
							Colors:      []Color{Green},
						},
						{
							PrintedName: "Káldring, la Varaescarcha",
							Colors:      []Color{Black, Blue},
						},
					},
					PrintedName: "",
				},
			},
		},
		{
			name: "missing opening array bracket",
			Input: `
				  {
				    "id": "00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "oracle_id": "2afbaa9a-c171-4a8b-90f3-5250d8498356",
				    "name": "Shardless Agent",
				    "printed_name": "Agente sin fragmento",
				    "lang": "es",
				    "uri": "https://api.scryfall.com/cards/00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "scryfall_uri": "https://scryfall.com/card/mh2/321/es/agente-sin-fragmento-(shardless-agent)?utm_source=api",
				    "cmc": 3.0,
				    "type_line": "Artifact Creature — Human Rogue",
				    "colors": [
						"G",
						"U"
				    ],
				    "color_identity": [
						"G",
						"U"
				    ],
				    "games": [
						"paper",
						"mtgo"
				    ],
				    "set_id": "c1c7eb8c-f205-40ab-a609-767cb296544e",
				    "set": "mh2",
				    "set_name": "Modern Horizons 2",
				    "collector_number": "321",
				    "rarity": "rare"
				  }
				]`,
			ThrowsError: true,
		},
		{
			name: "extra comma before closing array bracket",
			Input: `
				[
				  {
				    "id": "00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "oracle_id": "2afbaa9a-c171-4a8b-90f3-5250d8498356",
				    "name": "Shardless Agent",
				    "printed_name": "Agente sin fragmento",
				    "lang": "es",
				    "uri": "https://api.scryfall.com/cards/00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "scryfall_uri": "https://scryfall.com/card/mh2/321/es/agente-sin-fragmento-(shardless-agent)?utm_source=api",
				    "cmc": 3.0,
				    "type_line": "Artifact Creature — Human Rogue",
				    "colors": [
						"G",
						"U"
				    ],
				    "color_identity": [
						"G",
						"U"
				    ],
				    "games": [
						"paper",
						"mtgo"
				    ],
				    "set_id": "c1c7eb8c-f205-40ab-a609-767cb296544e",
				    "set": "mh2",
				    "set_name": "Modern Horizons 2",
				    "collector_number": "321",
				    "rarity": "rare"
				  },
				]`,
			ThrowsError: true,
		},
		{
			name: "missing closing array bracket",
			Input: `
				[
				  {
				    "id": "00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "oracle_id": "2afbaa9a-c171-4a8b-90f3-5250d8498356",
				    "name": "Shardless Agent",
				    "printed_name": "Agente sin fragmento",
				    "lang": "es",
				    "uri": "https://api.scryfall.com/cards/00017e6d-bf93-4dcf-8751-f50aba77e2d2",
				    "scryfall_uri": "https://scryfall.com/card/mh2/321/es/agente-sin-fragmento-(shardless-agent)?utm_source=api",
				    "cmc": 3.0,
				    "type_line": "Artifact Creature — Human Rogue",
				    "colors": [
						"G",
						"U"
				    ],
				    "color_identity": [
						"G",
						"U"
				    ],
				    "games": [
						"paper",
						"mtgo"
				    ],
				    "set_id": "c1c7eb8c-f205-40ab-a609-767cb296544e",
				    "set": "mh2",
				    "set_name": "Modern Horizons 2",
				    "collector_number": "321",
				    "rarity": "rare"
				  }
				`,
			ThrowsError: true,
		},
	}

	decoder := jsonDecoder[ScryfallCard]{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cards, err := decoder.decodeArray(bytes.NewBufferString(test.Input))
			if test.ThrowsError {
				if err == nil {
					t.Fatalf(
						"test %s: decoding json array %s should have failed but did not",
						test.name,
						test.Input,
					)
				} else {
					return
				}
			}

			if err != nil {
				t.Fatalf(
					"test %s: decoding json array %s failed with error %v, but should have decoded correctly",
					test.name,
					test.Input,
					err,
				)
			}

			if !reflect.DeepEqual(cards, test.Expected) {
				t.Fatalf(
					"test %s: decoding json array %s should have resulted in cards %#v but got %#v instead",
					test.name,
					test.Input,
					test.Expected,
					cards,
				)
			}
		})
	}
}
