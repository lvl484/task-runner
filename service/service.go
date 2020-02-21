package service

import (
	"context"
	"fmt"

	"github.com/lvl484/task-runner/database"
	"github.com/lvl484/task-runner/model"
	"github.com/lvl484/task-runner/scheduler"
)

type Service struct {
	database  database.Database
	scheduler *scheduler.Scheduler
}

func NewService(database database.Database, scheduler *scheduler.Scheduler) *Service {
	return &Service{
		database:  database,
		scheduler: scheduler,
	}
}

func (s *Service) CreateTask(ctx context.Context, input *model.TaskInput) (string, error) {
	task := &model.Task{TaskInput: *input}
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

func (s *Service) CreateAction(ctx context.Context, input *model.TaskInput) (string, error) {
	task := &model.Task{TaskInput: *input, IsAction: true}
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
	err := s.scheduler.UnscheduleTask(id)
	if err != nil {
		return fmt.Errorf("unschedule: %w", err)
	}
	err = s.database.DeleteTask(ctx, id)
	if err != nil {
		return fmt.Errorf("database: %w", err)
	}
	return nil
}

func (s *Service) UpdateTask(ctx context.Context, id string, input *model.TaskInput) error {
	task := &model.Task{TaskInput: *input}
	err := s.scheduler.UnscheduleTask(id)
	if err != nil {
		return fmt.Errorf("unschedule: %w", err)
	}
	err = s.database.UpdateTask(ctx, id, task)
	if err != nil {
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
