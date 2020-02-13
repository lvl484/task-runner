package database

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/lvl484/task-runner/model"
	"sync"
)

type Memory struct {
	tasks map[string]*model.Task
	mu    *sync.RWMutex // TODO: add usage Mutex to methods
}

func NewMemory() *Memory {
	return &Memory{
		tasks: make(map[string]*model.Task),
		mu:    new(sync.RWMutex),
	}
}

func (m *Memory) CreateTask(ctx context.Context, task *model.Task) (string, error) {
	task.ID = uuid.New().String()
	m.tasks[task.ID] = task
	return task.ID, nil
}

func (m *Memory) DeleteTask(ctx context.Context, id string) error {
	delete(m.tasks, id)
	return nil
}

func (m *Memory) UpdateTask(ctx context.Context, id string, task *model.Task) error {
	task.ID = id
	m.tasks[id] = task
	return nil
}

func (m *Memory) GetTask(ctx context.Context, id string) (*model.Task, error) {
	if task, ok := m.tasks[id]; ok {
		return task, nil
	}
	return nil, fmt.Errorf("task %s not found", id)
}
