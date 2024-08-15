package database

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
)

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type IApplicationRepository interface {
	Save(ctx context.Context, application domain.Application) error
	SearchApplications(ctx context.Context, applicationIDs []uuid.UUID) (applications []domain.Application, err error)
}

type ApplicationRepository struct {
	DB *sqlx.DB
}

func NewApplicationRepository(db *sqlx.DB) IApplicationRepository {
	return &ApplicationRepository{
		DB: db,
	}
}

func (repository *ApplicationRepository) Save(ctx context.Context, application domain.Application) error {
	applicationDB := ToApplicationDB(application)

	query := sq.Insert("applications").Columns(
		"id",
		"external_id",
		"tax_id",
	).Values(applicationDB.ID, applicationDB.ExternalID, applicationDB.TaxID).RunWith(repository.DB)

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = repository.DB.ExecContext(ctx, s, args)

	return err
}

func (repository *ApplicationRepository) SearchApplications(ctx context.Context, applicationIDs []uuid.UUID) (applications []domain.Application, err error) {
	query := sq.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From("applications")

	if len(applicationIDs) > 0 {
		query = query.Where(squirrel.Eq{"id": applicationIDs})
	}

	query.RunWith(repository.DB)

	s, args, err := query.ToSql()
	if err != nil {
		return applications, err
	}

	var applicationsDB []ApplcationDB
	err = repository.DB.SelectContext(ctx, &applicationsDB, s, args...)
	if err != nil {
		return applications, err
	}

	for _, applicationDB := range applicationsDB {
		application := FromApplicationDB(applicationDB)
		applications = append(applications, *application)
	}

	return applications, err
}
