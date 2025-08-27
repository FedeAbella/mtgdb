-- name: InsertCards :copyfrom
INSERT INTO cards (
	scryfall_id,
	set_id,
	name,
	collector_number,
	color_identity,
	colors,
	language_code,
	spanish_name,
	rarity,
	type_line,
	scryfall_api_uri,
	scryfall_web_uri,
	scryfall_oracle_id,
	created_at,
	updated_at
) VALUES (
    $1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9,
	$10,
	$11,
	$12,
	$13,
	$14,
	$15
);
