-- name: UpdateSet :exec
UPDATE sets
SET  code = $2,
    name = $3,
    updated_at = $4
WHERE scryfall_id = $1;
