package actions

import (
	"context"
	"fmt"

	"github.com/lvl484/task-runner/model"
)

// Select predefined action
func SelectActions(ctx context.Context, task *model.Task) (string, error) {
	action := task.Script
	switch action {
	case "CurrentTime":
		result := CurrentTime()
		return result, nil
	case "CurrentOS":
		result := CurrentOS()
		return result, nil
	case "CurrentCPU":
		result := CurrentCPU()
		return result, nil
	default:
		fmt.Println("Undefined action")
		return "Undefined action", nil
	}
}