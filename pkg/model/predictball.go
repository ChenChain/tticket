package model

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"
	"time"
	"tticket/pkg/dal"
)

type PredictBall struct {
	Ball
	PredictLotteryDrawingTime string

	Strategy string
	OrderNum int // 顺序号： 一天会产生多个号码，号码从高到低会有一个顺序号
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
			"num7",
			"strategy",
			"order_num"}), // column needed to be updated
	}).Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func FindPredictBalls(ctx context.Context, pageSize, pageNum int64) ([]*PredictBall, error) {
	res := make([]*PredictBall, 0)
	err := dal.DB.Model(&PredictBall{}).Order("id desc").Offset(int(pageSize * (pageNum - 1))).Limit(int(pageSize)).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindPredictBallByDrawingTime(times []string) ([]*PredictBall, error) {
	res := make([]*PredictBall, 0)
	if err := dal.DB.Model(&PredictBall{}).Where("predict_lottery_drawing_time in (?) ", times).Find(&res).Error; err != nil {
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

func GetPredictLotteryDrawingTime() string {
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
