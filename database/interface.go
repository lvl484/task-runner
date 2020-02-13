package database

import (
	"context"
	"github.com/lvl484/task-runner/model"
)

type Interface interface {
	CreateTask(ctx context.Context, task *model.Task) (string, error)
	DeleteTask(ctx context.Context, id string) error
	UpdateTask(ctx context.Context, id string, task *model.Task) error
	GetTask(ctx context.Context, id string) (*model.Task, error)
}
