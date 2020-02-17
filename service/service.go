package service

import (
	"context"
	"fmt"

	"github.com/lvl484/task-runner/database"
	"github.com/lvl484/task-runner/model"
	"github.com/lvl484/task-runner/scheduler"
)

type Service struct {
	database  database.Interface
	scheduler *scheduler.Scheduler
}

func NewService(database database.Interface, scheduler *scheduler.Scheduler) *Service{
	return &Service{
		database:  database,
		scheduler: scheduler,
	}
}

func (s *Service) CreateTask(ctx context.Context, task *model.Task) (string, error) {
	id, err := s.database.CreateTask(ctx, task)
	if err != nil {
		return "", fmt.Errorf("database: %w", err)
	}
	err = s.scheduler.ScheduleTask(ctx, task)
	if err != nil {
		return "", fmt.Errorf("schedule: %w", err)
	}
	return id, nil
}

func (s *Service) CreateAction(ctx context.Context, task *model.Task) (string, error) {
	task.IsAction = true
	id, err := s.database.CreateTask(ctx, task)
	if err != nil {
		return "", fmt.Errorf("database: %w", err)
	}
	err = s.scheduler.ScheduleTask(ctx, task)
	if err != nil {
		return "", fmt.Errorf("schedule: %w", err)
	}
	return id, nil
}

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	err := s.scheduler.UnscheduleTask()
	if err != nil {
		return fmt.Errorf("unschedule: %w", err)
	}
	err = s.database.DeleteTask(ctx, id)
	if err != nil {
		return fmt.Errorf("database: %w", err)
	}
	return nil
}

func (s *Service) UpdateTask(ctx context.Context, id string, task *model.Task) error {
	err := s.scheduler.UnscheduleTask()
	if err != nil {
		return fmt.Errorf("unschedule: %w", err)
	}
	err = s.database.UpdateTask(ctx, id, task)
	if err !=  nil {
		return fmt.Errorf("database: %w", err)
	}
	err = s.scheduler.ScheduleTask(ctx, task)
	if err != nil {
		return fmt.Errorf("schedule: %w", err)
	}
	return nil
}

func (s *Service) GetTask(ctx context.Context, id string) (*model.Task, error) {
	task, err := s.database.GetTask(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("database: %w", err)
	}
	return task, nil
}
