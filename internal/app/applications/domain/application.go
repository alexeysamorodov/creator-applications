package domain

import (
	"time"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	"github.com/google/uuid"
)

type Application struct {
	ID         uuid.UUID
	RequestID  int64
	Name       string
	Attributes []valueobjects.ApplicationAttribute
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewApplication(requestID int64, name string, attributes []valueobjects.ApplicationAttribute) *Application {
	now := time.Now().UTC()

	return &Application{
		ID:         uuid.New(),
		RequestID:  requestID,
		Name:       name,
		Attributes: attributes,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
