package domain

import (
	"time"

	"github.com/google/uuid"
)

type ChangeStateNotifyTask struct {
	Task
	OldState ApplicationState
	NewState ApplicationState
}

func CreateNewChangeStateNotifyTask(oldState ApplicationState, newState ApplicationState) *ChangeStateNotifyTask {
	return &ChangeStateNotifyTask{
		Task:     *CreateNewTask(TaskType_ChangeStateNotify),
		OldState: oldState,
		NewState: newState,
	}
}

func CreateChangeStateNotifyTaskFromDB(
	id uuid.UUID,
	taskType TaskType,
	state TaskState,
	createdAt time.Time,
	updatedAt time.Time,
	oldState ApplicationState,
	newState ApplicationState,
) *ChangeStateNotifyTask {
	return &ChangeStateNotifyTask{
		Task: *CreateTaskFromDB(
			id,
			taskType,
			state,
			createdAt,
			updatedAt,
		),
		OldState: oldState,
		NewState: newState,
	}
}
