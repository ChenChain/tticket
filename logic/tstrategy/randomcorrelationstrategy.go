package tstrategy

import (
	"context"
	"go.uber.org/zap"
	"math/rand"
	"time"
	"tticket/pkg/log"
	"tticket/pkg/model"
	"tticket/pkg/util"
)

/**
随机相关策略：
	在历史中奖数据中，随机选择

拿出最近100条双色球记录：
	arr[i][j]： i代表第i条记录， j代表第i条记录的第j个球， 其中 arr[i][6]为蓝球， 其他为红球

	predict[i]： 代表预测的球的号码， 0<=i<=6, predict[6]为蓝球， 其他为红球

	定义一个关联因子t： 0<=t<=1,  t代表了predict[i] 与 最近100条双色球记录的 关联程度

	举例：当t = 0 时， 那么predict[i]中的数字与 arr[][i]的数字毫无关联，需要尽量排除arr[][i]中的数字
	当 t = 1 时， predict[i]中的数字都从arr[][i]的数字选取
	当 t = 0.3时， predict[i]的数字有30%的概率从arr[][i]的数字选取

*/

type RandomCorrelationStrategy struct {
}

func NewRandomCorrelationStrategy() *RandomCorrelationStrategy {
	r := &RandomCorrelationStrategy{}
	return r
}

func (s *RandomCorrelationStrategy) Name() string {
	return "random correlation strategy"
}

func (s *RandomCorrelationStrategy) weight() int64 {
	return 100
}

func (s *RandomCorrelationStrategy) Predict(ctx context.Context) ([]int64, error) {
	balls, err := s.getBallData(ctx)
	if err != nil {
		return nil, err
	}
	arr1 := make([]int64, 0)
	arr2 := make([]int64, 0)
	arr3 := make([]int64, 0)
	arr4 := make([]int64, 0)
	arr5 := make([]int64, 0)
	arr6 := make([]int64, 0)
	arr7 := make([]int64, 0)

	predictArr := make([]int64, 7)
	for _, ball := range balls {
		arr1 = append(arr1, ball.Num1)
		arr2 = append(arr2, ball.Num2)
		arr3 = append(arr3, ball.Num3)
		arr4 = append(arr4, ball.Num4)
		arr5 = append(arr5, ball.Num5)
		arr6 = append(arr6, ball.Num6)
		arr7 = append(arr7, ball.Num7)
	}

	predictArr[0] = predictCorrelationNumber(predictFactor(), model.RED_BALL_TYPE, arr1, predictArr)
	predictArr[1] = predictCorrelationNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr2), predictArr)
	predictArr[2] = predictCorrelationNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr3), predictArr)
	predictArr[3] = predictCorrelationNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr4), predictArr)
	predictArr[4] = predictCorrelationNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr5), predictArr)
	predictArr[5] = predictCorrelationNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr6), predictArr)
	predictArr[6] = predictCorrelationNumber(predictFactor(), model.BLUE_BALL_TYPE, arr7, []int64{})
	return predictArr, nil
}

func (s *RandomCorrelationStrategy) ballDataNum() int64 {
	// Todo: config in yaml
	return 1000
}

func (s *RandomCorrelationStrategy) getBallData(ctx context.Context) ([]*model.Ball, error) {
	balls, err := model.FindBalls(ctx, s.ballDataNum(), 1)
	if err != nil {
		log.Error(ctx, "Failed to find balls", zap.Error(err))
		return nil, err
	}
	return balls, nil
}

// 生成预测因子
func predictFactor() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := r.Int63() % 101
	return float64(res / 100.0)
}

func predictCorrelationNumber(factor float64, ballType string, arr, predictArr []int64) int64 {
	var ballArr, otherArr []int64
	var res int64
	if ballType == model.BLUE_BALL_TYPE {
		ballArr = model.BLUE_BALL_NUMS
	} else {
		ballArr = model.RED_BALL_NUMS
	}
	otherArr = util.ArrayDifferentSet(ballArr, arr)
	if len(otherArr) == 0 {
		otherArr = util.ArrayDifferentSet(ballArr, predictArr)
	}

	nowPredictFactor := predictFactor()
	if nowPredictFactor < factor {
		// 从arr里面选
		res = util.RandomChooseOne(arr)
	}
	res = util.RandomChooseOne(otherArr)
	return res
}
