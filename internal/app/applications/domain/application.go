package domain

import "github.com/google/uuid"

type Application struct {
	ID        uuid.UUID
	RequestID int64
	TaxID     string
}

func NewApplication(requestID int64, taxID string) *Application {
	return &Application{
		ID:        uuid.New(),
		RequestID: requestID,
		TaxID:     taxID,
	}
}
