package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/lvl484/task-runner/database"
	"github.com/lvl484/task-runner/model"
	"github.com/lvl484/task-runner/runner"
)

type Scheduler struct {
	runner   runner.Runner
	database database.Database
	tasks    map[string]chan struct{}
	mu       *sync.Mutex
}

func NewScheduler(runner runner.Runner, database database.Database) (*Scheduler, error) {
	return &Scheduler{
		runner:   runner,
		database: database,
		tasks:    make(map[string]chan struct{}),
		mu:       new(sync.Mutex),
	}, nil
}

func cancelExecutions(duration time.Duration, done chan struct{}) {
	select {
	case <-done:
		return
	case <-time.After(duration):
		close(done)
	}
}

func (s *Scheduler) execution(done chan struct{}, task *model.Task) {
	defer fmt.Printf("[%s] Ticker stopped!\n", task.ID)

	time.Sleep(task.Schedule.UntilStartTime())

	ticker := time.NewTicker(task.Schedule.Interval.Duration())
	defer ticker.Stop()

	if !task.Schedule.EndAt.IsZero() {
		go cancelExecutions(task.Schedule.UntilEndTime(), done)
	}
	if task.Schedule.Count > 0 {
		go cancelExecutions(task.Schedule.UntilEndCount(), done)
	}

	var t = time.Now()
	for {
		err := s.run(task.ID)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("[%s] executed at = %v\n", task.ID, t)
		select {
		case <-done:
			fmt.Printf("[%s] Done!\n", task.ID)
			fmt.Printf("\nHistory was saved with [%s]\n", task.ID)
			return
		case t = <-ticker.C:
		}
	}
}

func (s *Scheduler) ScheduleTask(task *model.Task) error {
	fmt.Printf("ScheduleTask %s\n", task.ID)
	fmt.Printf("[%s] start at = %v\n", task.ID, task.Schedule.StartAt)
	fmt.Printf("[%s] end at = %v\n", task.ID, task.Schedule.EndAt)
	fmt.Printf("[%s] interval = %v\n", task.ID, task.Schedule.Interval)
	fmt.Printf("[%s] count = %d\n", task.ID, task.Schedule.Count)

	done := make(chan struct{})

	s.mu.Lock()
	s.tasks[task.ID] = done
	s.mu.Unlock()

	go s.execution(done, task)
	return nil
}

func (s *Scheduler) UnscheduleTask(id string) error {
	fmt.Printf("UnscheduleTask %s\n", id)
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if ok {
		close(task)
		delete(s.tasks, id)
	}
	return nil
}

// run() receive task id and getting task from database
// the set Status as running, and TimeStart as current time.
// Do updating task in database with new Status and TimeStart
// Execution task return Status Succeed if task was successful done otherwise Failed
// After that TimeFinish set as current time to record time when task was done
// Output set as result of task execution
// Do updating task in database with new Output and TimeSFinish
func (s *Scheduler) run(id string) error {
	var ctx = context.TODO()
	task, err := s.database.GetTask(ctx, id)
	if err != nil {
		return fmt.Errorf("get task: %w", err)
	}

	e := new(model.Execution)

	e.Status = model.Running
	e.StartTime = time.Now()
	err = s.database.UpdateTask(ctx, task.ID, task)
	if err != nil {
		return fmt.Errorf("update running task: %w", err)
	}
	output, err := s.runner.Execute(ctx, task)
	if err != nil {
		e.Status = model.Failed
		output = err.Error()
	} else {
		e.Status = model.Succeed
	}
	e.FinishTime = time.Now()
	e.Output = output
	task.Executions = append(task.Executions, *e)
	err = s.database.UpdateTask(ctx, task.ID, task)
	if err != nil {
		return fmt.Errorf("update completed task: %w", err)
	}
	return nil
}
