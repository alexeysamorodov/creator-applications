package domain

import (
	"time"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	"github.com/google/uuid"
)

type Application struct {
	ID         uuid.UUID
	RequestID  int64
	State      ApplicationState
	Name       string
	Attributes []valueobjects.ApplicationAttribute
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ApplicationState int32

const (
	ApplicationState_None    ApplicationState = 0
	ApplicationState_Created ApplicationState = 1
)

func NewApplication(requestID int64, name string, attributes []valueobjects.ApplicationAttribute) *Application {
	now := time.Now().UTC()

	return &Application{
		ID:         uuid.New(),
		RequestID:  requestID,
		State:      ApplicationState_Created,
		Name:       name,
		Attributes: attributes,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
