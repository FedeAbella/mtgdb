-- name: UpdateCard :exec
UPDATE cards
SET scryfall_oracle_id = $2,
    name = $3,
    updated_at = $4
WHERE id = $1;
