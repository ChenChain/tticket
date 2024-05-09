package tstrategy

import (
	"context"
	"tticket/pkg/model"
)

type Strategy interface {
	Name() string
	Predict(ctx context.Context) ([]int64, error) // 预测计算方法
	weight() int64                                // 权重 0～100
	getBallData(ctx context.Context) ([]*model.Ball, error)
	ballDataNum() int64 // 需要分析的数据数量
}

var strategies = map[string]Strategy{}

func Register(strategy Strategy) {
	if strategies[strategy.Name()] != nil {
		panic("duplicate strategy: " + strategy.Name())
	}
	strategies[strategy.Name()] = strategy
}

func Select() *Strategy {
	total := int64(0)
	for _, strategy := range strategies {
		total += strategy.weight()
	}
	// todo 权重随机选择
	return nil
}
