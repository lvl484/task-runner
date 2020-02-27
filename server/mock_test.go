package server

import (
	"context"
	"time"

	"github.com/lvl484/task-runner/model"
)

type mockService struct {
	err error
}

func (m mockService) CreateTask(ctx context.Context, input *model.TaskInput) (string, error) {
	return "CreateTask", m.err
}

func (m mockService) CreateAction(ctx context.Context, input *model.TaskInput) (string, error) {
	return "CreateAction", m.err
}

func (m mockService) DeleteTask(ctx context.Context, id string) error {
	return m.err
}

func (m mockService) UpdateTask(ctx context.Context, id string, input *model.TaskInput) error {
	return m.err
}

func (m mockService) UpdateAction(ctx context.Context, id string, input *model.TaskInput) error {
	return m.err
}

func (m mockService) GetTask(ctx context.Context, id string) (*model.Task, error) {
	return &model.Task{
		ID:       "GetTask",
		IsAction: true,
		TaskInput: model.TaskInput{
			Script: "CurrentTime",
			Schedule: model.Schedule{
				StartAt:  time.Date(2020, 2, 28, 15, 0, 0, 0, time.Local),
				EndAt:    time.Date(2020, 2, 28, 15, 20, 0, 0, time.Local),
				Count:    2,
				Interval: model.Duration(time.Second * 120),
			},
		},
		Executions: []model.Execution{
			{
				StartTime:  time.Date(2020, 2, 28, 15, 0, 0, 0, time.Local),
				FinishTime: time.Date(2020, 2, 28, 15, 10, 0, 0, time.Local),
				Status:     model.Running,
				Output:     "From MOCK",
			},
		},
	}, m.err
}
