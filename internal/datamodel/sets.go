package datamodel

import (
	"FedeAbella/mtgdb/internal/sqlc"
)

type Set struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (s *Set) Equals(dbSet sqlc.Set) bool {
	return s.Code == dbSet.Code && s.Name == dbSet.Name
}
