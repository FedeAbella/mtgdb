-- name: GetAllCards :many
SELECT
    c.*,
    s.code set_code,
    s.name set_name
FROM
    cards c
INNER JOIN sets s ON c.set_id = s.scryfall_id
ORDER BY set_code, set_name ASC;
