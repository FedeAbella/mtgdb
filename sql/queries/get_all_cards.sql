-- name: GetAllCards :many
SELECT
    *
FROM
    cards
ORDER BY name ASC;
