-- name: InsertSets :copyfrom
INSERT INTO sets (scryfall_id, code, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);
