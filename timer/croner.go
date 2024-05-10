package timer

import (
	"context"
	"errors"
	"fmt"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"time"
	"tticket/pkg/dal"
	"tticket/pkg/log"
	"tticket/pkg/model"
)

/**
一个简单的cron任务执行器
暂时支持 周 级别调度
自定义cron: * * *  --> minute: hour: week
场景举例：需要每周1，3，6 12:00发邮件，cronExpr =  0 12 1,3,6
*/

type Croner struct {
}

func (c *Croner) Name() string {
	return fmt.Sprintf("looper:%d", os.Getpid())
}

func (c *Croner) Execute(ctx context.Context, tasks []*model.Task) error {
	for _, task := range tasks {
		e := &Expr{}
		if err := e.parse(task.Cron); err != nil {
			log.Error(ctx, "failed to parse cron", zap.Any("task", task), zap.Error(err))
			continue
		}
		if !e.IsArrival() {
			continue
		}
		// 更新task
		result := dal.DB.Model(&model.Task{}).Where(&model.Task{
			Name:        task.Name,
			ExecuteTime: task.ExecuteTime, // 乐观锁
		}).Updates(
			&model.Task{
				Name:        task.Name,
				Executor:    c.Name(),
				ExecuteTime: time.Time{},
			},
		)

		if result.RowsAffected == 0 && result.Error == nil {
			continue
		}
		if result.Error != nil {
			log.Error(ctx, "failed to update task", zap.Any("task", task))
			return result.Error
		}

		go func() {
			log.Info(ctx, "start to execute task", zap.Any("task", task))
			err := taskMap[task.Name](ctx)
			if err != nil {
				log.Error(ctx, "failed to execute task", zap.String("task", task.Name))
				// need metrics
			}
		}()

	}
	return nil
}

type Expr struct {
	Minute  int64
	Hour    int64
	Weekday []int64
}

func (e *Expr) parse(expr string) error {
	arr := strings.Split(expr, " ")
	if len(arr) < 3 {
		return errors.New("invalid cron expr")
	}
	m, err := strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return err
	}
	h, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return err
	}
	weekdays := strings.Split(arr[2], ",")
	weekdayArr := make([]int64, 0)
	for _, v := range weekdays {
		w, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		weekdayArr = append(weekdayArr, w)
	}
	e.Minute = m
	e.Hour = h
	e.Weekday = weekdayArr
	return nil
}
func (e *Expr) IsArrival() bool {
	now := time.Now()
	w := now.Weekday()
	if w == 0 {
		w = 7
	}
	if !funk.ContainsInt64(e.Weekday, int64(w)) {
		return false
	}

	tmp := time.Date(now.Year(), now.Month(), now.Day(), int(e.Hour), int(e.Minute), 0, 0, now.Location())
	return !tmp.After(now)
}
