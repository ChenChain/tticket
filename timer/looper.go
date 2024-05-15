package timer

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"time"
	"tticket/pkg/dal"
	"tticket/pkg/log"
	"tticket/pkg/model"
)

// 秒级别定时任务

type Looper struct {
}

func (l *Looper) Name() string {
	return fmt.Sprintf("looper:%d", os.Getpid())
}

func (l *Looper) Execute(ctx context.Context, task *model.Task) error {
	if int64(time.Since(task.ExecuteTime).Seconds()) < task.IntervalSecond {
		return nil
	}

	// 更新task
	result := dal.DB.Model(&model.Task{}).Where(&model.Task{
		Name:        task.Name,
		ExecuteTime: task.ExecuteTime, // 乐观锁
	}).Updates(
		&model.Task{
			Name:        task.Name,
			Executor:    l.Name(),
			ExecuteTime: time.Now(),
		},
	)

	// 没更新到
	if result.RowsAffected == 0 && result.Error == nil {
		return nil
	}
	if result.Error != nil {
		log.Error(ctx, "failed to update task", zap.Any("task", task))
		return result.Error
	}
	log.Info(ctx, "start to execute task", zap.Any("task", task))
	f := taskMap[task.Name]
	if f == nil {
		log.Error(ctx, "execute task func is nil", zap.String("task", task.Name))
		// need metrics
		return errors.New("func is nil")
	}
	err := f(ctx)
	if err != nil {
		log.Error(ctx, "execute task err", zap.String("task", task.Name), zap.Error(err))
		// need metrics
		return err
	}
	return nil
}
