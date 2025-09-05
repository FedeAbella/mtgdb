-- +goose Up
CREATE TABLE cards (
    scryfall_id UUID PRIMARY KEY,
    set_id UUID NOT NULL REFERENCES sets(scryfall_id) ON DELETE CASCADE ,
    name TEXT NOT NULL,
    collector_number TEXT NOT NULL,
    color_identity TEXT,
    colors TEXT,
    language_code TEXT NOT NULL,
    spanish_name TEXT,
    rarity TEXT,
    type_line TEXT NOT NULL,
    scryfall_api_uri TEXT NOT NULL UNIQUE,
    scryfall_web_uri TEXT NOT NULL UNIQUE,
    scryfall_oracle_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE cards;
