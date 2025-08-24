-- name: InsertSets :copyfrom
INSERT INTO sets (code, name, created_at, updated_at) VALUES ($1, $2, $3, $4);
