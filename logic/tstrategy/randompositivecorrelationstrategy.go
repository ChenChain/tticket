package tstrategy

import (
	"context"
	"go.uber.org/zap"
	"tticket/pkg/log"
	"tticket/pkg/model"
	"tticket/pkg/util"
)

/**
随机正相关策略：
 预测结果随机产生，预测的数据均来源历史N天的中奖数据
*/

type RandomPositiveStrategy struct {
}

func NewRandomPositiveStrategy() *RandomPositiveStrategy {
	return &RandomPositiveStrategy{}
}

func (r RandomPositiveStrategy) Name() string {
	return "random positive strategy"
}

func (r RandomPositiveStrategy) Predict(ctx context.Context) ([]int64, error) {
	balls, err := r.getBallData(ctx)
	if err != nil {
		return nil, err
	}
	reds := balls[:len(model.RED_BALL_NUMS)]
	blues := balls[:len(model.BLUE_BALL_NUMS)]
	arr1 := make([]int64, 0)
	arr2 := make([]int64, 0)
	arr3 := make([]int64, 0)
	arr4 := make([]int64, 0)
	arr5 := make([]int64, 0)
	arr6 := make([]int64, 0)
	arr7 := make([]int64, 0)

	predictArr := make([]int64, 7)
	for _, ball := range reds {
		arr1 = append(arr1, ball.Num1)
		arr2 = append(arr2, ball.Num2)
		arr3 = append(arr3, ball.Num3)
		arr4 = append(arr4, ball.Num4)
		arr5 = append(arr5, ball.Num5)
		arr6 = append(arr6, ball.Num6)
	}

	for _, ball := range blues {
		arr7 = append(arr7, ball.Num7)
	}

	predictArr[0] = predictPositiveNumber(model.RED_BALL_TYPE, arr1, predictArr)
	predictArr[1] = predictPositiveNumber(model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr2), predictArr)
	predictArr[2] = predictPositiveNumber(model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr3), predictArr)
	predictArr[3] = predictPositiveNumber(model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr4), predictArr)
	predictArr[4] = predictPositiveNumber(model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr5), predictArr)
	predictArr[5] = predictPositiveNumber(model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr6), predictArr)
	predictArr[6] = predictPositiveNumber(model.BLUE_BALL_TYPE, arr7, []int64{})
	return predictArr, nil
}

func (r RandomPositiveStrategy) weight() int64 {
	return 100
}

func (r RandomPositiveStrategy) ballDataNum() int64 {
	// Todo: config in yaml
	return 100
}

func (r RandomPositiveStrategy) getBallData(ctx context.Context) ([]*model.Ball, error) {
	balls, err := model.FindBalls(ctx, r.ballDataNum(), 1)
	if err != nil {
		log.Error(ctx, "Failed to find balls", zap.Error(err))
		return nil, err
	}
	return balls, nil
}

func predictPositiveNumber(ballType string, arr, predictArr []int64) int64 {
	var ballArr, otherArr []int64
	if ballType == model.BLUE_BALL_TYPE {
		ballArr = model.BLUE_BALL_NUMS
	} else {
		ballArr = model.RED_BALL_NUMS
	}
	otherArr = util.ArrayDifferentSet(ballArr, arr)
	if len(otherArr) == 0 {
		otherArr = util.ArrayDifferentSet(ballArr, predictArr)
	}

	return util.RandomChooseOne(otherArr)
}
