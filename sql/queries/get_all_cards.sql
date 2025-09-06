-- name: GetAllCards :many
SELECT
    *
FROM
    cards c
ORDER BY name ASC;
