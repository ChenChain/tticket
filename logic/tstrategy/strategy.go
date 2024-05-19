package tstrategy

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"sort"
	"time"
	"tticket/pkg/log"
	"tticket/pkg/model"
)

type Strategy interface {
	Name() string
	Predict(ctx context.Context) ([]int64, error) // 预测计算方法
	weight() int64                                // 权重 0～100
}

var strategies = map[string]Strategy{}

func Register(strategy Strategy) {
	if strategies[strategy.Name()] != nil {
		panic("duplicate strategy: " + strategy.Name())
	}
	strategies[strategy.Name()] = strategy
}

func Init() {
	Register(NewRandomCorrelationStrategy())
	Register(NewRandomStrategy())
	Register(NewRandomPositiveStrategy())
	Register(NewRandomNegativeStrategy())
}

func Select() Strategy {
	total := int64(0)
	arr := make([]int64, len(strategies))
	strategyArr := make([]Strategy, len(strategies))
	i := 0
	for _, strategy := range strategies {
		total += strategy.weight()
		arr[i] = total
		strategyArr[i] = strategy
		i++
	}
	score := rand.Int63n(total)
	index := 0
	for i, v := range arr {
		if score < v {
			index = i
			break
		}
	}
	return strategyArr[index]
}

func PredictBall(ctx context.Context) error {
	strategy := Select()
	arr, err := strategy.Predict(ctx)
	if err != nil {
		log.Error(ctx, "failed to predict", zap.String("strategy", strategy.Name()), zap.Error(err))
		return err
	}
	tmp := arr[:6]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	arr = append(tmp, arr[6])

	ball := &model.PredictBall{
		Ball: model.Ball{
			Num1: arr[0],
			Num2: arr[1],
			Num3: arr[2],
			Num4: arr[3],
			Num5: arr[4],
			Num6: arr[5],
			Num7: arr[6],
		},
		// 当前日期的下一个周2， 4， 7
		PredictLotteryDrawingTime: getPredictLotteryDrawingTime(),
		Strategy:                  strategy.Name(),
	}
	for i := 0; i < 3; i++ {
		if err = model.CreatePredictBall(ctx, ball); err == nil {
			break
		}
	}
	return err
}

func getPredictLotteryDrawingTime() string {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekday := now.Weekday()
	delta := 0
	switch weekday {
	case time.Sunday:
	case time.Monday:
		delta = 1
	case time.Tuesday:
	case time.Wednesday:
		delta = 1
	case time.Thursday:
	case time.Friday:
		delta = 2
	case time.Saturday:
		delta = 1
	}
	date = date.AddDate(0, 0, delta)
	return fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
}
