-- +goose Up
CREATE TABLE set_cards (
    id UUID PRIMARY KEY,
    scryfall_id UUID NOT NULL UNIQUE,
    mtgjson_uuid UUID NOT NULL UNIQUE,
    mtgjson_v4_id UUID NOT NULL UNIQUE,
    set_code TEXT NOT NULL REFERENCES sets (code) ON DELETE CASCADE,
    card_id UUID NOT NULL REFERENCES cards (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE set_cards;
