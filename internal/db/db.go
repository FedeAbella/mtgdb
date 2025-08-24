package db

import (
	"FedeAbella/mtgdb/internal/sqlc"
)

type DbConf struct {
	Queries *sqlc.Queries
}
