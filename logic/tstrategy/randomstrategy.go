package tstrategy

import (
	"context"
	"tticket/pkg/model"
	"tticket/pkg/util"
)

/**
随机策略：
 预测结果随机产生，与历史中奖数据无关
*/

type RandomStrategy struct {
}

func NewRandomStrategy() *RandomStrategy {
	r := &RandomStrategy{}
	return r
}

func (r RandomStrategy) Name() string {
	return "random strategy"
}

func (r RandomStrategy) Predict(ctx context.Context) ([]int64, error) {
	predictArr := make([]int64, 7)

	predictArr[0] = randomPredictNumber(model.RED_BALL_TYPE, predictArr)
	predictArr[1] = randomPredictNumber(model.RED_BALL_TYPE, predictArr)
	predictArr[2] = randomPredictNumber(model.RED_BALL_TYPE, predictArr)
	predictArr[3] = randomPredictNumber(model.RED_BALL_TYPE, predictArr)
	predictArr[4] = randomPredictNumber(model.RED_BALL_TYPE, predictArr)
	predictArr[5] = randomPredictNumber(model.RED_BALL_TYPE, predictArr)
	predictArr[6] = randomPredictNumber(model.BLUE_BALL_TYPE, []int64{})
	return predictArr, nil
}

func (r RandomStrategy) weight() int64 {
	return 100
}

func randomPredictNumber(ballType string, arr []int64) int64 {
	var otherArr []int64
	if ballType == model.BLUE_BALL_TYPE {
		otherArr = util.ArrayDifferentSet(model.BLUE_BALL_NUMS, arr)
	} else {
		otherArr = util.ArrayDifferentSet(model.RED_BALL_NUMS, arr)
	}
	if len(otherArr) == 0 {
		otherArr = arr
	}
	return util.RandomChooseOne(otherArr)
}
