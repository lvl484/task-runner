package actions

import (
	"context"
	"errors"
	"fmt"

	"github.com/lvl484/task-runner/model"
)

const (
	PredefinedCurrentTime = "CurrentTime"
	PredefinedCurrentOS   = "CurrentOS"
	PredefinedCurrentCPU  = "CurrentCPU"
)

// SelectActions check json field task.Script. Compare it with all Predefined actions.
// After that run the matched action and return result of its execution or error if action not found
func SelectActions(ctx context.Context, task *model.Task) (string, error) {
	action := task.Script
	switch action {
	case PredefinedCurrentTime:
		return CurrentTime(), nil
	case PredefinedCurrentOS:
		return CurrentOS(), nil
	case PredefinedCurrentCPU:
		return CurrentCPU(), nil
	default:
		fmt.Println("\r\nUndefined action")
		return "", errors.New("undefined action")
	}
}