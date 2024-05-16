package timer

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
	"tticket/logic/tmail"
	"tticket/logic/tspider"
	"tticket/logic/tstrategy"
	"tticket/logic/tuser"
	"tticket/pkg/log"
	"tticket/pkg/model"
	"tticket/pkg/util"
)

type Executor interface {
	Name() string
	Execute(context.Context, *model.Task) error
}

var executorMap = map[model.TaskType]Executor{}
var taskMap = map[string]func(ctx context.Context) error{}

func Init(ctx context.Context) {
	RegisterExecutor(model.DEFER_TASK_TYPE, &Defer{})
	RegisterExecutor(model.LOOP_TASK_TYPE, &Looper{})
	RegisterExecutor(model.CRON_TASK_TYPE, &Croner{})

	RegisterTask("cache_user", tuser.CacheUserMail)
	RegisterTask("predict_ball", tstrategy.PredictBall)
	RegisterTask("spider_lottery", tspider.Logic)
	RegisterTask("send_mail", tmail.Send)

	go start(ctx)
}

func RegisterTask(taskName string, taskFunc func(ctx context.Context) error) {
	if taskMap[taskName] != nil {
		panic(fmt.Sprintf("duplicate task:%s", taskName))
	}
	taskMap[taskName] = taskFunc
}

func RegisterExecutor(taskType model.TaskType, executor Executor) {
	if executorMap[taskType] != nil {
		panic(fmt.Sprintf("duplicate taskType:%d", taskType))
	}
	executorMap[taskType] = executor
}

func start(ctx context.Context) {
	t := time.NewTicker(10 * time.Second)
	for range t.C {
		execute(ctx)
	}
}

func execute(ctx context.Context) {
	// get all task from db
	tasks, err := model.FindTask(ctx, 0)
	if err != nil {
		log.Error(ctx, "failed to find loop task", zap.Error(err))
		return
	}
	model.FillEventID(tasks)

	taskArrMap := make(map[model.TaskType][]*model.Task)
	for _, t := range tasks {
		taskType := t.Type
		arr := taskArrMap[taskType]
		if arr == nil {
			arr = make([]*model.Task, 0)
		}
		arr = append(arr, t)
		taskArrMap[taskType] = arr
	}

	for k, v := range taskArrMap {
		executor := executorMap[k]
		if executor == nil {
			log.Error(ctx, "the executor is nil", zap.Any("task_type", k))
			continue
		}
		for _, task := range v {
			go func() {
				tmpCtx := context.WithValue(ctx, util.TASK_EVENT_ID, task.EventID)
				err := executor.Execute(tmpCtx, task)
				if err != nil {
					log.Error(ctx, "execute task err", zap.Any("task", v))
					// should metrics
				}
			}()
		}
	}
}
