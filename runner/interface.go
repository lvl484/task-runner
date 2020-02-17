package runner

import (
	"context"

	"github.com/lvl484/task-runner/model"
)

type Interface interface {
	Execute(ctx context.Context, task *model.Task) (string, error)
}
