package model

import (
	"context"
	"time"
	"tticket/pkg/dal"
)

type PredictBall struct {
	Ball
	PredictLotteryDrawingTime string
	CreatedTime               time.Time
	UpdatedTime               time.Time
	Strategy                  string
}

func FindPredictBalls(ctx context.Context, pageSize, pageNum int64) ([]*PredictBall, error) {
	res := make([]*PredictBall, 0)
	err := dal.DB.Model(&PredictBall{}).Offset(int(pageSize * (pageNum - 1))).Limit(int(pageSize)).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindPredictBallByDrawingTime(times []string) ([]*PredictBall, error) {
	res := make([]*PredictBall, 0)
	if err := dal.DB.Model(&PredictBall{}).Where("predict_lottery_drawing_time in (?) ", times).Find(res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GenerateErrorPredictBall() *PredictBall {
	return &PredictBall{
		Ball: Ball{
			LotteryDrawingTime: "error",
			Num1:               -1,
		},
		PredictLotteryDrawingTime: "error",
		Strategy:                  "error",
	}
}
