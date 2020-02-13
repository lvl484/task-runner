package runner

import (
	"context"
	"fmt"
	"github.com/lvl484/task-runner/model"
	"os/exec"
)

type Bash struct {

}

func NewBash() *Bash{
	return &Bash{}
}

func (b *Bash) Execute(ctx context.Context, task *model.Task) (string, error) {
	cmd := exec.Command("bash", "-c", task.Script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w\n\n%s", err, string(output))
	}
	return string(output), nil
}
