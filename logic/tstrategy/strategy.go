package tstrategy

import (
	"context"
	"go.uber.org/zap"
	"math/rand"
	"tticket/pkg/conf"
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

func Select() []Strategy {
	nums := conf.Config.Strategy.UserStrategyNum
	if nums > len(strategies) {
		nums = len(strategies)
	}
	res := make([]Strategy, 0)
	tmp := map[string]Strategy{}
	for k, v := range strategies {
		tmp[k] = v
	}

	for k := 0; k < nums; k++ {
		total := int64(0)
		arr := make([]int64, len(tmp))
		strategyArr := make([]Strategy, len(tmp))
		i := 0
		for _, strategy := range tmp {
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
		res = append(res, strategyArr[index])
		delete(tmp, strategyArr[index].Name())
	}
	return res
}

func PredictBall(ctx context.Context) error {
	strategies := Select()
	nextTime := model.GetPredictLotteryDrawingTime()
	for index, strategy := range strategies {
		arr, err := strategy.Predict(ctx)
		if err != nil {
			log.Error(ctx, "failed to predict", zap.String("strategy", strategy.Name()), zap.Error(err))
			return err
		}

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
			PredictLotteryDrawingTime: nextTime,
			Strategy:                  strategy.Name(),
			OrderNum:                  index + 1,
		}
		for i := 0; i < 3; i++ {
			if err = model.CreatePredictBall(ctx, ball); err == nil {
				break
			}
		}
	}
	return nil
}
