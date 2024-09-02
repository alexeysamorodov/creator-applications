package database

import (
	"context"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
)

type MockApplicationRepository struct {
	Applications map[uuid.UUID]ApplcationDB
	DB           sqlx.DB
}

func NewMockApplicationRepository() IApplicationRepository {
	return &MockApplicationRepository{
		Applications: make(map[uuid.UUID]ApplcationDB),
	}
}

func (repository *MockApplicationRepository) Save(ctx context.Context, application domain.Application) error {
	applicationDB, err := ToApplicationDB(application)

	if err != nil {
		return err
	}

	repository.Applications[applicationDB.ID] = *applicationDB

	return nil
}

func (repository *MockApplicationRepository) SearchApplications(ctx context.Context, applicationIDs []uuid.UUID) (applications []domain.Application, err error) {
	for _, id := range applicationIDs {
		applicationDB, ok := repository.Applications[id]
		if ok {
			application, err := FromApplicationDB(applicationDB)
			if err != nil {
				return applications, err
			}

			applications = append(applications, *application)
		}
	}

	return applications, nil
}
