package model

import "time"

type Task struct {
	ID         string    `json:"id"`
	Script     string    `json:"script"`
	TimeStart  time.Time `json:"time_start"`
	TimeFinish time.Time `json:"time_finish"`
	Status     Status    `json:"status"`
	Output     string    `json:"output"`
}
