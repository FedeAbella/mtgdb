package source

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/sqlc"
)

func Test_EqualsSqlcSet(t *testing.T) {
	tests := []struct {
		name    string
		set     Set
		sqlcSet sqlc.Set
		expect  bool
	}{
		{
			name: "all fields equal",
			set: Set{
				Code:       "apc",
				Name:       "Apocalypse",
				ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			},
			sqlcSet: sqlc.Set{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Code: "apc",
				Name: "Apocalypse",
			},
			expect: true,
		},
		{
			name: "different scryfall id",
			set: Set{
				Code:       "apc",
				Name:       "Apocalypse",
				ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			},
			sqlcSet: sqlc.Set{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce1"),
					Valid: true,
				},
				Code: "apc",
				Name: "Apocalypse",
			},
			expect: false,
		},
		{
			name: "different code",
			set: Set{
				Code:       "apc",
				Name:       "Apocalypse",
				ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			},
			sqlcSet: sqlc.Set{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Code: "ust",
				Name: "Apocalypse",
			},
			expect: false,
		},
		{
			name: "different name",
			set: Set{
				Code:       "apc",
				Name:       "Apocalypse",
				ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			},
			sqlcSet: sqlc.Set{
				ScryfallID: pgtype.UUID{
					Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
					Valid: true,
				},
				Code: "apc",
				Name: "Planeshift",
			},
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.set.Equals(&test.sqlcSet); got != test.expect {
				t.Fatalf("test %s expected %v but got %v", test.name, test.expect, got)
			}
		})
	}
}

func Test_ToDbInsertSet(t *testing.T) {
	now := time.Date(2025, 9, 5, 20, 48, 0, 0, time.UTC)
	set := Set{
		Code:       "apc",
		Name:       "Apocalypse",
		ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
	}
	want := sqlc.InsertSetsParams{
		ScryfallID: pgtype.UUID{
			Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			Valid: true,
		},
		Code: "apc",
		Name: "Apocalypse",
		CreatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
	}

	t.Run("map set into sqlc insert params", func(t *testing.T) {
		got := set.ToDbInsertSet(now)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected %#v, got %#v", want, got)
		}
	})
}

func Test_ToDbUpdateSet(t *testing.T) {
	now := time.Date(2025, 9, 5, 20, 48, 0, 0, time.UTC)
	set := Set{
		Code:       "apc",
		Name:       "Apocalypse",
		ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
	}
	want := sqlc.UpdateSetParams{
		ScryfallID: pgtype.UUID{
			Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
			Valid: true,
		},
		Code: "apc",
		Name: "Apocalypse",
		UpdatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
	}

	t.Run("map set into sqlc insert params", func(t *testing.T) {
		got := set.ToDbUpdateSet(now)
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected %#v, got %#v", want, got)
		}
	})
}
