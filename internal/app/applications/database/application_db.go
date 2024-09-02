package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ApplcationDB struct {
	ID         uuid.UUID       `db:"id"`
	ExternalID int64           `db:"external_id"`
	Data       *sql.NullString `db:"data"`
	CreatedAt  time.Time       `db:"created_at"`
	UpdatedAt  time.Time       `db:"updated_at"`
}

type ApplicationDataDB struct {
	Name       string                   `json:"name"`
	Attributes []ApplicationAttributeDB `json:"attributes"`
}

type ApplicationAttributeDB struct {
	Name string `json:"name"`
}
