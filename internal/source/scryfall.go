package source

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/google/uuid"
)

type LanguageCode = string
type Color = string
type Rarity = string

const (
	English   LanguageCode = "en"
	Spanish   LanguageCode = "es"
	GamePaper string       = "paper"

	White Color = "W"
	Blue  Color = "U"
	Black Color = "B"
	Red   Color = "R"
	Green Color = "G"

	Common   = "common"
	Uncommon = "uncommon"
	Rare     = "rare"
	Special  = "special"
	Mythic   = "mythic"
	Bonus    = "bonus"
)

type ScryfallCardFace struct {
	Colors      []string `json:"colors"`
	PrintedName string   `json:"printed_name"`
}

type ScryfallCard struct {
	CMC              float32            `json:"cmc"`
	CollectorNumber  string             `json:"collector_number"`
	ColorIdentity    []Color            `json:"color_identity"`
	Colors           []Color            `json:"colors"`
	Faces            []ScryfallCardFace `json:"card_faces"`
	Games            []string           `json:"games"`
	LanguageCode     LanguageCode       `json:"lang"`
	Name             string             `json:"name"`
	PrintedName      string             `json:"printed_name"`
	Rarity           Rarity             `json:"rarity"`
	ScryfallAPIURI   string             `json:"uri"`
	ScryfallId       uuid.UUID          `json:"id"`
	ScryfallOracleId uuid.UUID          `json:"oracle_id"`
	ScryfallSetId    uuid.UUID          `json:"set_id"`
	ScryfallWebURI   string             `json:"scryfall_uri"`
	SetCode          string             `json:"set"`
	SetName          string             `json:"set_name"`
	TypeLine         string             `json:"type_line"`
}

func (sfCard *ScryfallCard) unpack() (Set, CardPrinting) {
	return Set{
			Code:       sfCard.SetCode,
			Name:       sfCard.SetName,
			ScryfallId: sfCard.ScryfallSetId,
		}, CardPrinting{
			CollectorNumber:  sfCard.CollectorNumber,
			ColorIdentity:    strings.Join(sfCard.ColorIdentity, ""),
			Colors:           sfCard.getColors(),
			Language:         string(sfCard.LanguageCode),
			Name:             sfCard.Name,
			NameSPA:          sfCard.getSpanishName(),
			Rarity:           sfCard.Rarity,
			ScryfallAPIURI:   sfCard.ScryfallAPIURI,
			ScryfallId:       sfCard.ScryfallId,
			ScryfallOracleId: sfCard.ScryfallOracleId,
			ScryfallWebURI:   sfCard.ScryfallWebURI,
			SetScryfallId:    sfCard.ScryfallSetId,
			TypeLine:         sfCard.TypeLine,
		}
}

func (sfCard *ScryfallCard) getColors() string {
	if len(sfCard.Colors) > 0 {
		return strings.Join(sfCard.Colors, "")
	}

	if len(sfCard.Faces) == 0 {
		return ""
	}

	faceColors := make([]Color, 0)
	faceColors = append(faceColors, sfCard.Faces[0].Colors...)
	faceColors = append(faceColors, sfCard.Faces[1].Colors...)
	if len(faceColors) == 0 {
		return ""
	}

	sort.Strings(faceColors)
	return strings.Join(slices.Compact(faceColors), "")
}

func (sfCard *ScryfallCard) getSpanishName() string {
	if sfCard.LanguageCode == English {
		return ""
	}

	if sfCard.PrintedName != "" {
		return sfCard.PrintedName
	}

	if len(sfCard.Faces) == 0 || sfCard.Faces[0].PrintedName == "" {
		return ""
	}

	return fmt.Sprintf("%s // %s", sfCard.Faces[0].PrintedName, sfCard.Faces[1].PrintedName)
}

func scryfallToSetsCards(
	sfCards []ScryfallCard,
) (map[uuid.UUID]Set, map[uuid.UUID]CardPrinting) {
	setMap := make(map[uuid.UUID]Set)
	printMap := make(map[uuid.UUID]CardPrinting)

	for _, sfCard := range sfCards {
		if !slices.Contains(sfCard.Games, GamePaper) {
			continue
		}

		if sfCard.LanguageCode != English && sfCard.LanguageCode != Spanish {
			continue
		}

		set, printing := sfCard.unpack()
		setMap[set.ScryfallId] = set
		printMap[printing.ScryfallId] = printing
	}

	return setMap, printMap
}
