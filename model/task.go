package model

import (
	"time"
)

type TaskInput struct {
	Script   string   `json:"script"`
	Schedule Schedule `json:"schedule"`
}

type Task struct {
	ID         string    `json:"id"`
	TimeStart  time.Time `json:"time_start"`
	TimeFinish time.Time `json:"time_finish"`
	Status     Status    `json:"status"`
	Output     string    `json:"output"`
	IsAction   bool      `json:"is_action"`
	TaskInput  `json:"task_input"`
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
