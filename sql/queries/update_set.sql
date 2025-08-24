-- name: UpdateSet :exec
UPDATE sets
SET name = $2,
    updated_at = $3
WHERE code = $1;
