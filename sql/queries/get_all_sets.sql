-- name: GetAllSets :many
SELECT
    *
FROM
    sets
ORDER BY code ASC;
