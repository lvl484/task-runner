package runner

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/lvl484/task-runner/actions"
	"github.com/lvl484/task-runner/model"
)

type Bash struct {

}

func NewBash() *Bash {
	return &Bash{}
}

func (b *Bash) Execute(ctx context.Context, task *model.Task) (string, error) {
	// checks if the task is action
	if task.IsAction {
		return actions.SelectActions(ctx, task)
	}

	cmd := exec.Command("bash", "-c", task.Script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w\n\n%s", err, string(output))
	}
	return string(output), nil
}