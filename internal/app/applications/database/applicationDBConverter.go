package database

import (
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
)

func ToApplicationDB(application domain.Application) *ApplcationDB {
	result := ApplcationDB{
		ID:         application.ID,
		ExternalID: application.RequestID,
	}

	return &result
}

func FromApplicationDB(application ApplcationDB) *domain.Application {
	result := domain.Application{
		ID:        application.ID,
		RequestID: application.ExternalID,
	}

	return &result
}
