package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/lvl484/task-runner/model"
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

// CreateTask receive input task, generate new UUID for it.
// Set that task to map with all tasks and return task ID
func (m *Memory) CreateTask(ctx context.Context, task *model.Task) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task.ID = uuid.New().String()

	m.tasks[task.ID] = task
	return task.ID, nil
}

// DeleteTask receive task ID. Do deleting required task from map by ID
func (m *Memory) DeleteTask(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.tasks, id)
	return nil
}

// UpdateTask receive task ID, which need to update,
// and task, which must replace the selected by ID task
func (m *Memory) UpdateTask(ctx context.Context, id string, task *model.Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	task.ID = id
	m.tasks[id] = task
	return nil
}

// GetTask receive task ID, which need to get.
// If selected task exists return task or error if task not found
func (m *Memory) GetTask(ctx context.Context, id string) (*model.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, ok := m.tasks[id]
	if ok {
		return task, nil
	}
	return nil, fmt.Errorf("task %s not found", id)
}
