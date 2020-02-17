package runner

import (
	"context"
	"fmt"
	"github.com/lvl484/task-runner/actions"
	"github.com/lvl484/task-runner/model"
	"os/exec"
)

const (
	LayoutTimeFormat = "02 Jan 06 15:04 -0700"
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