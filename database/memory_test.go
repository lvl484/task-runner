package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lvl484/task-runner/model"
)

func TestMemory_CreateTask(t *testing.T) {
	m := NewMemory()
	m.tasks = make(map[string]*model.Task)
	task := &model.Task{
		ID:       "CreateTask",
		IsAction: false,
		TaskInput: model.TaskInput{
			Script: "pwd",
			Schedule: model.Schedule{
				StartAt:  time.Date(2020, 2, 28, 16, 30, 0, 0, time.Local),
				EndAt:    time.Date(2020, 2, 28, 16, 50, 0, 0, time.Local),
				Count:    2,
				Interval: model.Duration(time.Minute * 10),
			},
		},
		Executions: nil,
	}

	tests := []struct {
		want  map[string]*model.Task
		input *model.Task
		msg   string
	}{
		{
			want:  m.tasks,
			input: task,
			msg:   "case 1",
		},
	}

	for _, tt := range tests {
		ctx := context.Background()

		id, err := m.CreateTask(ctx, tt.input)
		if err != nil {
			t.Fatal(err)
		}

		receivedTask, err := m.GetTask(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("CREATE:", id)
		fmt.Println("GET:", receivedTask.ID)
		equalsIDs := id == receivedTask.ID
		if !equalsIDs {
			t.Errorf("want %v got %v", task, m.tasks[id])
		}
	}
}

func TestMemory_DeleteTaskTask(t *testing.T) {
	m := NewMemory()
	m.tasks = make(map[string]*model.Task)
	m.tasks["CreateTask"] = &model.Task{
		ID:       "CreateTask",
		IsAction: false,
		TaskInput: model.TaskInput{
			Script: "pwd",
			Schedule: model.Schedule{
				StartAt:  time.Date(2020, 2, 28, 16, 30, 0, 0, time.Local),
				EndAt:    time.Date(2020, 2, 28, 16, 50, 0, 0, time.Local),
				Count:    2,
				Interval: model.Duration(time.Minute * 10),
			},
		},
		Executions: nil,
	}

	tests := []struct {
		want  int
		input string
		msg   string
	}{
		{
			want:  0,
			input: "CreateTask",
			msg:   "case 1",
		},
	}

	for _, tt := range tests {
		ctx := context.Background()

		err := m.DeleteTask(ctx, tt.input)
		if err != nil {
			t.Fatal(err)
		}

		if len(m.tasks) != tt.want {
			t.Errorf("want %v got %v", tt.want, len(m.tasks))
		}
	}
}

func TestMemory_UpdateTask(t *testing.T) {
	m := NewMemory()
	m.tasks = make(map[string]*model.Task)
	m.tasks["CreateTask"] = &model.Task{
		ID:       "CreateTask",
		IsAction: false,
		TaskInput: model.TaskInput{
			Script: "pwd",
			Schedule: model.Schedule{
				StartAt:  time.Date(2020, 2, 28, 16, 30, 0, 0, time.Local),
				EndAt:    time.Date(2020, 2, 28, 16, 50, 0, 0, time.Local),
				Count:    2,
				Interval: model.Duration(time.Minute * 10),
			},
		},
		Executions: nil,
	}

	task := &model.Task{
		ID:       "CreateTask",
		IsAction: false,
		TaskInput: model.TaskInput{
			Script: "echo Hello!",
			Schedule: model.Schedule{
				StartAt:  time.Date(2020, 2, 28, 16, 30, 0, 0, time.Local),
				EndAt:    time.Date(2020, 2, 28, 16, 50, 0, 0, time.Local),
				Count:    2,
				Interval: model.Duration(time.Minute * 10),
			},
		},
		Executions: nil,
	}

	tests := []struct {
		want  map[string]*model.Task
		input *model.Task
		msg   string
	}{
		{
			want:  m.tasks,
			input: task,
			msg:   "case 1",
		},
	}

	for _, tt := range tests {
		ctx := context.Background()

		err := m.UpdateTask(ctx, "CreateTask", tt.input)
		if err != nil {
			t.Fatal(err)
		}

		receivedTask, err := m.GetTask(ctx, "CreateTask")
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println("CREATE:", task.Script)
		fmt.Println("GET:", receivedTask.Script)

		equalsIDs := task.Script == receivedTask.Script
		if !equalsIDs {
			t.Errorf("want %v got %v", "echo Hello!", m.tasks["CreateTask"].Script)
		}
	}
}

func TestMemory_GetTask(t *testing.T) {
	m := NewMemory()
	m.tasks = make(map[string]*model.Task)
	m.tasks["CreateTask"] = &model.Task{
		ID:       "CreateTask",
		IsAction: false,
		TaskInput: model.TaskInput{
			Script: "pwd",
			Schedule: model.Schedule{
				StartAt:  time.Date(2020, 2, 28, 16, 30, 0, 0, time.Local),
				EndAt:    time.Date(2020, 2, 28, 16, 50, 0, 0, time.Local),
				Count:    2,
				Interval: model.Duration(time.Minute * 10),
			},
		},
		Executions: nil,
	}

	tests := []struct {
		want  map[string]*model.Task
		input string
		msg   string
	}{
		{
			want:  m.tasks,
			input: "Unknown ID",
			msg:   "case 1",
		},
	}

	for _, tt := range tests {
		ctx := context.Background()

		_, err := m.GetTask(ctx, tt.input)
		if err == nil {
			t.Errorf("want err got nil")
		}
	}
}
