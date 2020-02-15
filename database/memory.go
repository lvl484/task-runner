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
	mu    *sync.RWMutex
}

func NewMemory() *Memory {
	return &Memory{
		tasks: make(map[string]*model.Task),
		mu:    new(sync.RWMutex),
	}
}

func (m *Memory) CreateTask(ctx context.Context, task *model.Task) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task.ID = uuid.New().String()
	m.tasks[task.ID] = task
	return task.ID, nil
}

func (m *Memory) DeleteTask(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.tasks, id)
	return nil
}

func (m *Memory) UpdateTask(ctx context.Context, id string, task *model.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	task.ID = id
	m.tasks[id] = task
	return nil
}

func (m *Memory) GetTask(ctx context.Context, id string) (*model.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, ok := m.tasks[id]
	if ok {
		return task, nil
	}
	return nil, fmt.Errorf("task %s not found", id)
}
