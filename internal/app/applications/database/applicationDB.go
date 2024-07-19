package database

import "github.com/google/uuid"

type ApplcationDB struct {
	ID        uuid.UUID
	TaxID     string
	RequestID int64
}
