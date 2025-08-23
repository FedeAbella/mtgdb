-- name: GetAllSets :many
SELECT
    *
FROM
    sets
ORDER BY name ASC;
