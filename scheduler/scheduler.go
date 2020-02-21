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
	runner   runner.Runner
	database database.Database
}

func NewScheduler(ctx context.Context, runner runner.Runner, database database.Database) (*Scheduler, error) {
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

// run() receive task id and getting task from database
// the set Status as running, and TimeStart as current time.
// Do updating task in database with new Status and TimeStart
// Execution task return Status Succeed if task was successful done otherwise Failed
// After that TimeFinish set as current time to record time when task was done
// Output set as result of task execution
// Do updating task in database with new Output and TimeSFinish
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
