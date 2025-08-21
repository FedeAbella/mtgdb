#!/bin/bash

source .env
(
    cd sql/schema || exit 1
    goose postgres $DB_URL down && goose postgres $DB_URL up
)
