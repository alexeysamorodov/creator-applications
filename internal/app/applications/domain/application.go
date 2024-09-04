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
	Tasks      []ITask
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// TODO: add field HasNotCompletedTasks
}

type ApplicationState int32

const (
	ApplicationState_None    ApplicationState = 0
	ApplicationState_Created ApplicationState = 1
)

func CreateNewApplication(
	requestID int64,
	name string,
	attributes []valueobjects.ApplicationAttribute,
) *Application {
	now := time.Now().UTC()
	var appTasks []ITask
	appTasks = append(appTasks, *CreateNewChangeStateNotifyTask(ApplicationState_None, ApplicationState_Created))

	return &Application{
		ID:         uuid.New(),
		RequestID:  requestID,
		State:      ApplicationState_Created,
		Name:       name,
		Attributes: attributes,
		Tasks:      appTasks,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func CreateApplicationFromDB(
	id uuid.UUID,
	requestID int64,
	state ApplicationState,
	name string,
	attributes []valueobjects.ApplicationAttribute,
	tasks []ITask,
	createdAt time.Time,
	updatedAt time.Time,
) *Application {
	return &Application{
		ID:         id,
		RequestID:  requestID,
		State:      state,
		Name:       name,
		Attributes: attributes,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}
