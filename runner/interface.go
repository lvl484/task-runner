package runner

import (
	"context"

	"github.com/lvl484/task-runner/model"
)

type Runner interface {
	Execute(ctx context.Context, task *model.Task) (string, error)
}
