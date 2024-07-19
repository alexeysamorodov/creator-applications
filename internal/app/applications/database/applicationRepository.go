package database

import (
	"context"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"

	"github.com/google/uuid"
)

type IApplicationRepository interface {
	Save(ctx context.Context, application domain.Application)
	SearchApplications(ctx context.Context, applicationIDs []uuid.UUID) (applications []domain.Application)
}

type ApplicationRepository struct {
	Applications map[uuid.UUID]ApplcationDB
}

func NewApplicationRepository() IApplicationRepository {
	return &ApplicationRepository{
		Applications: make(map[uuid.UUID]ApplcationDB),
	}
}

func (repository *ApplicationRepository) Save(ctx context.Context, application domain.Application) {
	applicationDB := ToApplicationDB(application)

	repository.Applications[applicationDB.ID] = *applicationDB
}

func (repository *ApplicationRepository) SearchApplications(ctx context.Context, applicationIDs []uuid.UUID) (applications []domain.Application) {
	for _, id := range applicationIDs {
		applicationDB, ok := repository.Applications[id]
		if ok {
			application := FromApplicationDB(applicationDB)

			applications = append(applications, *application)
		}
	}

	return applications
}
