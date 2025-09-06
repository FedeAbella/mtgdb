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

func Test_MapSetsToInsertAndUpdate(t *testing.T) {
	now := time.Date(2025, 9, 5, 21, 36, 0, 0, time.UTC)
	past := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name               string
		setsInDb           []sqlc.Set
		setsInFile         map[uuid.UUID]source.Set
		expectedInsertSets []sqlc.InsertSetsParams
		expectedUpdateSets []sqlc.UpdateSetParams
	}{
		{
			name:     "no sets in DB",
			setsInDb: []sqlc.Set{},
			setsInFile: map[uuid.UUID]source.Set{
				uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"): {
					Code:       "ori",
					Name:       "Magic Origins",
					ScryfallId: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
				},
				uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"): {
					Code:       "apc",
					Name:       "Apocalypse",
					ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				},
				uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"): {
					Code:       "moc",
					Name:       "March of the Machine Commander",
					ScryfallId: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
				},
				uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"): {
					Code:       "cmm",
					Name:       "Commander Masters",
					ScryfallId: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
				},
			},
			expectedInsertSets: []sqlc.InsertSetsParams{
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
						Valid: true,
					},
					Code: "ori",
					Name: "Magic Origins",
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
				},
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
						Valid: true,
					},
					Code: "moc",
					Name: "March of the Machine Commander",
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
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
						Valid: true,
					},
					Code: "cmm",
					Name: "Commander Masters",
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
			expectedUpdateSets: []sqlc.UpdateSetParams{},
		},
		{
			name: "all sets in DB",
			setsInDb: []sqlc.Set{
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
						Valid: true,
					},
					Code: "ori",
					Name: "Magic Origins",
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
				},
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
						Valid: true,
					},
					Code: "moc",
					Name: "March of the Machine Commander",
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
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
						Valid: true,
					},
					Code: "cmm",
					Name: "Commander Masters",
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
			setsInFile: map[uuid.UUID]source.Set{
				uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"): {
					Code:       "ori",
					Name:       "Magic Origins",
					ScryfallId: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
				},
				uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"): {
					Code:       "apc",
					Name:       "Apocalypse",
					ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				},
				uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"): {
					Code:       "moc",
					Name:       "March of the Machine Commander",
					ScryfallId: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
				},
				uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"): {
					Code:       "cmm",
					Name:       "Commander Masters",
					ScryfallId: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
				},
			},
			expectedInsertSets: []sqlc.InsertSetsParams{},
			expectedUpdateSets: []sqlc.UpdateSetParams{},
		},
		{
			name: "some sets to insert, some to update",
			setsInDb: []sqlc.Set{
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
						Valid: true,
					},
					Code: "ori",
					Name: "Magic Origins",
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
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
						Valid: true,
					},
					Code: "apc",
					Name: "Apocalypse",
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
			setsInFile: map[uuid.UUID]source.Set{
				uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"): {
					Code:       "ori",
					Name:       "Magic Origins -- Updated",
					ScryfallId: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
				},
				uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"): {
					Code:       "apc",
					Name:       "Apocalypse -- Updated",
					ScryfallId: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
				},
				uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"): {
					Code:       "moc",
					Name:       "March of the Machine Commander",
					ScryfallId: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
				},
				uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"): {
					Code:       "cmm",
					Name:       "Commander Masters",
					ScryfallId: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
				},
			},
			expectedInsertSets: []sqlc.InsertSetsParams{
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("6bba5de9-5afb-42af-a7eb-24ac854bf671"),
						Valid: true,
					},
					Code: "moc",
					Name: "March of the Machine Commander",
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
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("cd05036f-2698-43e6-a48e-5c8d82f0a551"),
						Valid: true,
					},
					Code: "cmm",
					Name: "Commander Masters",
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
			expectedUpdateSets: []sqlc.UpdateSetParams{
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("0eeb9a9a-20ac-404d-b55f-aeb7a43a7f62"),
						Valid: true,
					},
					Code: "ori",
					Name: "Magic Origins -- Updated",
					UpdatedAt: pgtype.Timestamp{
						Time:  now,
						Valid: true,
					},
				},
				{
					ScryfallID: pgtype.UUID{
						Bytes: uuid.MustParse("e4e00913-d08d-4899-86ea-5cf631e09ce0"),
						Valid: true,
					},
					Code: "apc",
					Name: "Apocalypse -- Updated",
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
			gotInsert, gotUpdate := mapSetsToInsertAndUpdate(test.setsInFile, test.setsInDb, now)
			for _, expectedInsert := range test.expectedInsertSets {
				if !slices.Contains(gotInsert, expectedInsert) {
					t.Fatalf(
						"test %s expected set %#v to be inserted, but got %#v",
						test.name,
						expectedInsert,
						gotInsert,
					)
				}
			}

			for _, expectedUpdate := range test.expectedUpdateSets {
				if !slices.Contains(gotUpdate, expectedUpdate) {
					t.Fatalf(
						"test %s expected set %#v to be updated, but got %#v",
						test.name,
						expectedUpdate,
						gotUpdate,
					)
				}
			}

		})
	}
}
