-- name: InsertCards :copyfrom
INSERT INTO cards (id, name, scryfall_oracle_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);
