package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ApplcationDB struct {
	ID         uuid.UUID      `db:"id"`
	ExternalID int64          `db:"external_id"`
	DataJson   sql.NullString `db:"data"`
	TasksJson  sql.NullString `db:"tasks"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
}

type ApplicationDataDB struct {
	State      string                   `json:"state"`
	Name       string                   `json:"name"`
	Attributes []ApplicationAttributeDB `json:"attributes"`
}

type ApplicationAttributeDB struct {
	Name string `json:"name"`
}

type TasksContainerDB struct {
	Tasks []TaskDataDB `json:"tasks"`
}

type TaskDataDB struct {
	Type     string `json:"type"`
	DataJson string `json:"data"`
}
