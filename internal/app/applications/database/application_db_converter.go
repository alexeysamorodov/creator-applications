package database

import (
	"database/sql"
	"encoding/json"

	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects"
	"github.com/alexeysamorodov/creator-applications/internal/app/applications/domain/valueobjects/errors"
)

func ConvertApplicationToApplicationDB(application *domain.Application) (*ApplcationDB, error) {
	appDataJson, err := convertApplicationToApplicationDataJson(application)
	if err != nil {
		return nil, err
	}

	tasksJson, err := convertTasksToTasksJson(application.Tasks)
	if err != nil {
		return nil, err
	}

	result := ApplcationDB{
		ID:         application.ID,
		ExternalID: application.RequestID,
		DataJson:   *appDataJson,
		TasksJson:  *tasksJson,
		CreatedAt:  application.CreatedAt,
		UpdatedAt:  application.UpdatedAt,
	}

	return &result, nil
}

func convertApplicationToApplicationDataJson(application *domain.Application) (*sql.NullString, error) {
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

	return &dataNullString, nil
}

func convertTasksToTasksJson(tasks []domain.ITask) (*sql.NullString, error) {
	var taskDataDBs []TaskDataDB
	for _, task := range tasks {
		taskDataDB, err := convertTaskToTaskDataDB(task)
		if err != nil {
			return nil, err
		}

		taskDataDBs = append(taskDataDBs, *taskDataDB)
	}

	tasksContainerDB := TasksContainerDB{
		Tasks: taskDataDBs,
	}

	dataBytes, err := json.Marshal(tasksContainerDB)
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

	return &dataNullString, nil
}

func convertTaskToTaskDataDB(task domain.ITask) (*TaskDataDB, error) {
	dataBytes, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}

	result := &TaskDataDB{
		Type:     task.GetTypeAsStr(),
		DataJson: string(dataBytes),
	}

	return result, nil
}

func ConvertApplicationDBToApplication(applicationDB *ApplcationDB) (*domain.Application, error) {
	var state domain.ApplicationState
	var name string
	var attributes []valueobjects.ApplicationAttribute
	var tasks []domain.ITask

	if applicationDB.DataJson.Valid {
		var appData ApplicationDataDB
		err := json.Unmarshal([]byte(applicationDB.DataJson.String), &appData)
		if err != nil {
			return nil, err
		}

		state, err = FromApplicationStateDb(appData.State)
		if err != nil {
			return nil, err
		}

		name = appData.Name
		attributes = FromAttributesDb(appData.Attributes)
	}

	if applicationDB.TasksJson.Valid {
		var tasksContainer TasksContainerDB
		err := json.Unmarshal([]byte(applicationDB.TasksJson.String), &tasksContainer)
		if err != nil {
			return nil, err
		}

		for _, taskDB := range tasksContainer.Tasks {
			task, err := ConvertFromTaskDataDBToTask(&taskDB)
			if err != nil {
				return nil, err
			}

			tasks = append(tasks, *task)
		}
	}

	result := domain.CreateApplicationFromDB(
		applicationDB.ID,
		applicationDB.ExternalID,
		state,
		name,
		attributes,
		tasks,
		applicationDB.CreatedAt,
		applicationDB.UpdatedAt,
	)

	return result, nil
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
	case domain.ApplicationState_None:
		result = "None"
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
		case "None":
			result = domain.ApplicationState_None
		case "Created":
			result = domain.ApplicationState_Created
		default:
			return domain.ApplicationState_None, errors.NewEnumStrOutOfRangeError("domain.ApplicationState", applicationState)
		}
	} else {
		result = domain.ApplicationState_Created
	}

	return result, nil
}

func ConvertFromTaskDataDBToTask(taskData *TaskDataDB) (*domain.ITask, error) {
	var result domain.ITask
	var err error

	switch taskData.Type {
	case domain.TaskTypeStr[domain.TaskType_ChangeStateNotify]:
		result, err = deserializeTaskDataDB[domain.ChangeStateNotifyTask](taskData.DataJson)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.NewEnumStrOutOfRangeError("domain.TaskType", taskData.Type)
	}

	return &result, nil
}

func deserializeTaskDataDB[T domain.ITask](dataJson string) (*T, error) {
	var task T
	err := json.Unmarshal([]byte(dataJson), &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
