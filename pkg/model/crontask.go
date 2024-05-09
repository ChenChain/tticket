package model

import (
	"context"
	"time"
	"tticket/pkg/dal"
)

// 秒级别定时任务

type TaskType int64

const DEFER_TASK_TYPE = 1 // 延迟任务
const LOOP_TASK_TYPE = 1  // 定时任务

type CronTask struct {
	ID             int64
	Name           string
	IntervalSecond int64
	Type           TaskType

	Executor    string    // 执行该任务的人
	ExecuteTime time.Time // 执行时间

	CreatedTime time.Time
	UpdatedTime time.Time
	DeletedTime time.Time
}

func FindLoopTask(ctx context.Context) ([]*CronTask, error) {
	// 目前只处理loop task
	res := make([]*CronTask, 0)
	if err := dal.DB.Model(&CronTask{Type: LOOP_TASK_TYPE}).Find(res).Error; err != nil {
		return nil, err
	}
}
