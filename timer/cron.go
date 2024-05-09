package timer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
	"tticket/logic/tuser"
	"tticket/pkg/dal"
	"tticket/pkg/log"
	"tticket/pkg/model"
)

// 秒级别定时任务

type Cron struct {
	Name string
}

var cron *Cron

var taskMap map[string]func(ctx context.Context) error

func Init(ctx context.Context) {
	cron = &Cron{}
	cron.Name = uuid.New().String()

	RegisterTask("cache_user", tuser.CacheUserMail)
	RegisterTask("spider_lottery", tuser.CacheUserMail)
	RegisterTask("send_mail", tuser.CacheUserMail)

	go start(ctx)
}

func RegisterTask(taskName string, taskFunc func(ctx context.Context) error) {
	if taskMap[taskName] != nil {
		panic(fmt.Sprintf("duplicate task:%s", taskName))
	}
	taskMap[taskName] = taskFunc
}

func start(ctx context.Context) {
	t := time.NewTicker(10 * time.Second)
	for range t.C {
		execute(ctx)
	}
}

func execute(ctx context.Context) {
	// get all task from db
	res, err := model.FindLoopTask(ctx)
	if err != nil {
		log.Error(ctx, "failed to find loop task", zap.Error(err))
		return
	}

	for _, task := range res {
		if int64(time.Since(task.ExecuteTime).Seconds()) < task.IntervalSecond {
			continue
		}

		// 更新task
		result := dal.DB.Model(&model.CronTask{}).Where(&model.CronTask{
			Name:        task.Name,
			ExecuteTime: task.ExecuteTime, // 乐观锁
		}).
			Updates(
				&model.CronTask{
					Name:        task.Name,
					Executor:    cron.Name,
					ExecuteTime: time.Time{},
				},
			)

		if result.RowsAffected == 0 && result.Error == nil {
			continue
		}
		if result.Error != nil {
			log.Error(ctx, "failed to update task", zap.Any("task", task))
			continue
		}
		go func() {
			log.Info(ctx, "start to execute task", zap.Any("task", task))
			err := taskMap[task.Name](ctx)
			if err != nil {
				log.Error(ctx, "failed to execute task", zap.String("task_name", task.Name))
				// need metrics
			}
		}()
	}
}
