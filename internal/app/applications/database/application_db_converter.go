package database

import (
	"database/sql"
	"encoding/json"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects/errors"
)

func ToApplicationDB(application *domain.Application) (*ApplcationDB, error) {
	state, err := ToApplicationStateDb(application.State)
	if err != nil {
		return nil, err
	}

	appDataDB := ApplicationDataDB{
		Name:       application.Name,
		State:      state,
		Attributes: ToAttributesDb(application.Attributes),
	}

	dataBytes, err := json.Marshal(appDataDB)
	if err != nil {
		return nil, err
	}

	var dataNullString sql.NullString
	if len(dataBytes) > 0 {
		dataNullString = sql.NullString{
			String: string(dataBytes),
			Valid:  true,
		}
	} else {
		dataNullString = sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	result := ApplcationDB{
		ID:         application.ID,
		ExternalID: application.RequestID,
		Data:       dataNullString,
		CreatedAt:  application.CreatedAt,
		UpdatedAt:  application.UpdatedAt,
	}

	return &result, nil
}

func ToDomain(applicationDB *ApplcationDB) (*domain.Application, error) {
	result := domain.Application{
		ID:        applicationDB.ID,
		RequestID: applicationDB.ExternalID,
		CreatedAt: applicationDB.CreatedAt,
		UpdatedAt: applicationDB.UpdatedAt,
	}

	if applicationDB.Data.Valid {
		var appData ApplicationDataDB
		err := json.Unmarshal([]byte(applicationDB.Data.String), &appData)
		if err != nil {
			return nil, err
		}

		state, err := FromApplicationStateDb(appData.State)
		if err != nil {
			return nil, err
		}

		result.State = state
		result.Name = appData.Name
		result.Attributes = FromAttributesDb(appData.Attributes)
	}

	return &result, nil
}

func ToAttributesDb(attributes []valueobjects.ApplicationAttribute) []ApplicationAttributeDB {
	var result []ApplicationAttributeDB

	for _, attribute := range attributes {
		attributeDB := ApplicationAttributeDB{
			Name: attribute.Name,
		}

		result = append(result, attributeDB)
	}

	return result
}

func FromAttributesDb(attributes []ApplicationAttributeDB) []valueobjects.ApplicationAttribute {
	var result []valueobjects.ApplicationAttribute

	for _, attributeDB := range attributes {
		attribute := valueobjects.ApplicationAttribute{
			Name: attributeDB.Name,
		}

		result = append(result, attribute)
	}

	return result
}

func ToApplicationStateDb(applicationState domain.ApplicationState) (string, error) {
	var result string

	switch applicationState {
	case domain.ApplicationState_Created:
		result = "Created"
	default:
		return "", errors.NewEnumOutOfRangeError(applicationState)
	}

	return result, nil
}

func FromApplicationStateDb(applicationState string) (domain.ApplicationState, error) {
	var result domain.ApplicationState
	if applicationState != "" {
		switch applicationState {
		case "Created":
			result = domain.ApplicationState_Created
		default:
			return domain.ApplicationState_None, errors.NewEnumOutOfRangeError(applicationState)
		}
	} else {
		result = domain.ApplicationState_Created
	}

	return result, nil
}
