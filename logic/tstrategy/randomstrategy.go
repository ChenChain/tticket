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
随机策略：

拿出最近100条双色球记录：
	arr[i][j]： i代表第i条记录， j代表第i条记录的第j个球， 其中 arr[i][6]为蓝球， 其他为红球

	predict[i]： 代表预测的球的号码， 0<=i<=6, predict[6]为蓝球， 其他为红球

	定义一个关联因子t： 0<=t<=1,  t代表了predict[i] 与 最近100条双色球记录的 关联程度

	举例：当t = 0 时， 那么predict[i]中的数字与 arr[][i]的数字毫无关联，需要尽量排除arr[][i]中的数字
	当 t = 1 时， predict[i]中的数字都从arr[][i]的数字选取
	当 t = 0.3时， predict[i]的数字有30%的概率从arr[][i]的数字选取

*/

type RandomStrategy struct {
	predictFactor float64
}

func NewRandomStrategy() *RandomStrategy {
	r := &RandomStrategy{}
	r.predictFactor = predictFactor()
	return r
}

func (s *RandomStrategy) Name() string {
	return "random strategy"
}

func (s *RandomStrategy) weight() int64 {
	return 100
}

func (s *RandomStrategy) Predict(ctx context.Context) ([]int64, error) {
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

	predictArr := make([]int64, 0)
	for _, ball := range balls {
		arr1 = append(arr1, ball.Num1)
		arr2 = append(arr2, ball.Num2)
		arr3 = append(arr3, ball.Num3)
		arr4 = append(arr4, ball.Num4)
		arr5 = append(arr5, ball.Num5)
		arr6 = append(arr6, ball.Num6)
		arr7 = append(arr7, ball.Num7)
	}

	predictArr[0] = predictNumber(predictFactor(), model.RED_BALL_TYPE, arr1)
	predictArr[1] = predictNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr2))
	predictArr[2] = predictNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr3))
	predictArr[3] = predictNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr4))
	predictArr[4] = predictNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr5))
	predictArr[5] = predictNumber(predictFactor(), model.RED_BALL_TYPE, util.ArrayAdd(predictArr, arr6))
	predictArr[6] = predictNumber(predictFactor(), model.BLUE_BALL_TYPE, arr7)

	return predictArr, nil
}

func (s *RandomStrategy) ballDataNum() int64 {
	// Todo: config in yaml
	return 100
}

func (s *RandomStrategy) getBallData(ctx context.Context) ([]*model.Ball, error) {
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

func predictNumber(factor float64, ballType string, arr []int64) int64 {
	var otherArr []int64
	if ballType == model.BLUE_BALL_TYPE {
		otherArr = util.ArrayDifferentSet(model.BLUE_BALL_NUMS, arr)
	} else {
		otherArr = util.ArrayDifferentSet(model.RED_BALL_NUMS, arr)
	}
	if len(otherArr) == 0 {
		otherArr = arr
	}
	nowPredictFactor := predictFactor()
	if nowPredictFactor < factor {
		// 从arr里面选
		return util.RandomChooseOne(arr)
	}
	return util.RandomChooseOne(otherArr)
}
