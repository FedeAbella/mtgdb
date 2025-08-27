-- name: UpdateCard :exec
UPDATE cards
SET set_id = $2,
    name = $3,
    collector_number = $4,
    color_identity = $5,
    colors = $6,
    language_code = $7,
    spanish_name = $8,
    rarity = $9,
    type_line = $10,
    scryfall_api_uri = $11,
    scryfall_web_uri = $12,
    scryfall_oracle_id = $13,
    updated_at = $14
WHERE scryfall_id = $1;
