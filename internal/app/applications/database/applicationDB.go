package database

import "github.com/google/uuid"

type ApplcationDB struct {
	ID         uuid.UUID `db:"id"`
	TaxID      string    `db:"tax_id"`
	ExternalID int64     `db:"external_id"`
}
