package domain

import "github.com/google/uuid"

type Application struct {
	ID        uuid.UUID
	RequestID int64
}

func NewApplication(requestID int64) *Application {
	return &Application{
		ID:        uuid.New(),
		RequestID: requestID,
	}
}
