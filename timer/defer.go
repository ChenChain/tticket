package timer

import (
	"context"
	"fmt"
	"os"
	"tticket/pkg/model"
)

type Defer struct {
}

func (d *Defer) Name() string {
	return fmt.Sprintf("defer:%d", os.Getpid())
}

func (d *Defer) Execute(ctx context.Context, task *model.Task) error {
	// TODO 延迟任务执行器
	return nil
}
