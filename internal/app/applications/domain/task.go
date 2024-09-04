package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	Type      TaskType
	State     TaskState
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ITask interface {
	GetID() uuid.UUID
	GetType() TaskType
	GetTypeAsStr() string
}

func (t Task) GetID() uuid.UUID {
	return t.ID
}

func (t Task) GetType() TaskType {
	return t.Type
}

func (t Task) GetTypeAsStr() string {
	return TaskTypeStr[t.Type]
}

type TaskType int32

const (
	TaskType_None              TaskType = 0
	TaskType_ChangeStateNotify TaskType = 1
)

var TaskTypeStr map[TaskType]string = map[TaskType]string{
	TaskType_ChangeStateNotify: "ChangeStateNotify",
}

type TaskState int32

const (
	TaskState_None      TaskState = 0
	TaskState_InProcess TaskState = 1
	TaskState_Completed TaskState = 2
)

func CreateNewTask(taskType TaskType) *Task {
	now := time.Now().UTC()
	id := uuid.New()
	return &Task{
		ID:        id,
		Type:      taskType,
		State:     TaskState_InProcess,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func CreateTaskFromDB(
	id uuid.UUID,
	taskType TaskType,
	state TaskState,
	createdAt time.Time,
	updatedAt time.Time,
) *Task {
	return &Task{
		ID:        id,
		Type:      taskType,
		State:     state,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
