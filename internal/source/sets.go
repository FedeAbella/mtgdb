package source

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"FedeAbella/mtgdb/internal/sqlc"
)

type Set struct {
	Code       string
	Name       string
	ScryfallId uuid.UUID
}

func (s *Set) Equals(dbSet *sqlc.Set) bool {
	return s.ScryfallId == dbSet.ScryfallID.Bytes && s.Code == dbSet.Code && s.Name == dbSet.Name
}

func (s *Set) ToDbInsertSet() sqlc.InsertSetsParams {
	return sqlc.InsertSetsParams{
		ScryfallID: pgtype.UUID{
			Bytes: s.ScryfallId,
			Valid: true,
		},
		Code: s.Code,
		Name: s.Name,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}
}

func (s *Set) ToDbUpdateSet() sqlc.UpdateSetParams {
	return sqlc.UpdateSetParams{
		ScryfallID: pgtype.UUID{
			Bytes: s.ScryfallId,
			Valid: true,
		},
		Code: s.Code,
		Name: s.Name,
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}
}
