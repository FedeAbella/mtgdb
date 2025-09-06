package db

import (
	"github.com/jackc/pgx/v5"

	"FedeAbella/mtgdb/internal/sqlc"
)

type DbConf struct {
	Conn    *pgx.Conn
	Queries *sqlc.Queries
}
