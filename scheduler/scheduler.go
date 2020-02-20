package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/lvl484/task-runner/database"
	"github.com/lvl484/task-runner/model"
	"github.com/lvl484/task-runner/runner"
)

type Scheduler struct {
	runner   runner.Interface
	database database.Interface
}

func NewScheduler(ctx context.Context, runner runner.Interface, database database.Interface) (*Scheduler, error) {
	return &Scheduler{
		runner:   runner,
		database: database,
	}, nil
}

func (s *Scheduler) ScheduleTask(ctx context.Context, task *model.Task) error {
	return s.run(ctx, task.ID)
}

func (s *Scheduler) UnscheduleTask() error {

	return nil
}

func (s *Scheduler) run(ctx context.Context, id string) error {
	task, err := s.database.GetTask(ctx, id)
	if err != nil {
		return fmt.Errorf("get task: %w", err)
	}
	task.Status = model.Running
	task.TimeStart = time.Now()
	err = s.database.UpdateTask(ctx, task.ID, task)
	if err != nil {
		return fmt.Errorf("update running task: %w", err)
	}
	output, err := s.runner.Execute(ctx, task)
	if err != nil {
		task.Status = model.Failed
		output = err.Error()
	} else {
		task.Status = model.Succeed
	}
	task.TimeFinish = time.Now()
	task.Output = output
	err = s.database.UpdateTask(ctx, task.ID, task)
	if err != nil {
		return fmt.Errorf("update completed task: %w", err)
	}
	return nil
}
