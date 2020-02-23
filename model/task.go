package model

import (
	"time"
)

type TaskInput struct {
	Script   string   `json:"script"`
	Schedule Schedule `json:"schedule"`
}

type Task struct {
	ID         string `json:"id"`
	IsAction   bool   `json:"is_action"`
	TaskInput  `json:"task_input"`
	Executions []Execution `json:"executions"`
}

type Execution struct {
	StartTime  time.Time `json:"start_time"`
	FinishTime time.Time `json:"finish_time"`
	Status     Status    `json:"status"`
	Output     string    `json:"output"`
}

func NewTask(input *TaskInput) *Task {
	input.Schedule.SetDefaultValue()
	return &Task{
		TaskInput: *input,
	}
}

func NewAction(input *TaskInput) *Task {
	input.Schedule.SetDefaultValue()
	return &Task{
		TaskInput: *input,
		IsAction:  true,
	}
}
