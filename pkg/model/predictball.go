package model

import (
	"context"
	"gorm.io/gorm/clause"
	"tticket/pkg/dal"
)

type PredictBall struct {
	Ball
	PredictLotteryDrawingTime string

	Strategy string
}

func (b *PredictBall) TableName() string {
	return "predict_ball"
}

func CreatePredictBall(ctx context.Context, m *PredictBall) error {
	err := dal.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "predict_lottery_drawing_time"}}, // key colum
		DoUpdates: clause.AssignmentColumns([]string{
			"num1",
			"num2",
			"num3",
			"num4",
			"num5",
			"num6",
			"num7"}), // column needed to be updated
	}).Create(&m).Error
	if err != nil {
		return err
	}
	return nil
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
			Num1: -1,
		},
		PredictLotteryDrawingTime: "1900-01-01",
		Strategy:                  "error",
	}
}
