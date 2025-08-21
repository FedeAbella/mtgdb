-- +goose Up
CREATE TABLE cards (
    id UUID PRIMARY KEY,
    scryfall_oracle_id UUID NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE cards;
