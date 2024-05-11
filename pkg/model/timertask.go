package model

import (
	"context"
	"github.com/google/uuid"
	"time"
	"tticket/pkg/dal"
)

// 10秒级别定时任务

type TaskType int64

const UNKNOWN_TASK_TYPE = 0
const DEFER_TASK_TYPE = 1 // 延迟任务
const LOOP_TASK_TYPE = 2  // 定时任务
const CRON_TASK_TYPE = 3  // cron定时任务

type Task struct {
	ID             int64
	EventID        string // log追踪
	Name           string
	IntervalSecond int64
	Type           TaskType

	Executor    string    // 执行该任务的人
	ExecuteTime time.Time // 执行时间
	Cron        string    // cron表达式

	CreatedTime time.Time
	UpdatedTime time.Time
	DeletedTime time.Time
}

func FindTask(ctx context.Context, ty TaskType) ([]*Task, error) {
	res := make([]*Task, 0)
	if err := dal.DB.Model(&Task{Type: ty}).Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FillEventID(tasks []*Task) {
	for _, t := range tasks {
		t.EventID = uuid.New().String()
	}
}
