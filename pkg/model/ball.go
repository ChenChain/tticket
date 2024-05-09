package model

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"time"
	"tticket/pkg/dal"
	"tticket/pkg/log"
	"tticket/pkg/util"
)

var RED_BALL_NUMS = []int64{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
}

var BLUE_BALL_NUMS = []int64{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
}

var RED_BALL_TYPE = "red"
var BLUE_BALL_TYPE = "blue"

type Ball struct {
	ID                 int64
	LotteryDrawingTime string // 开奖日期
	Num1               int64
	Num2               int64
	Num3               int64
	Num4               int64
	Num5               int64
	Num6               int64
	Num7               int64
	CreatedTime        time.Time
	UpdatedTime        time.Time
}

func (ball *Ball) GetBallNumsString() string {
	return fmt.Sprintf("%d %d %d %d %d %d %d",
		ball.Num1, ball.Num2, ball.Num3, ball.Num4, ball.Num5, ball.Num6, ball.Num7,
	)
}

func (ball *Ball) GetBallNumsArray() []int64 {
	return []int64{
		ball.Num1, ball.Num2, ball.Num3, ball.Num4, ball.Num5, ball.Num6, ball.Num7,
	}
}

func (ball *Ball) IsWinning(otherBall *Ball) bool {
	if ball.Num7 == otherBall.Num7 {
		return true
	}
	intersect := util.ArrayIntersect(ball.GetBallNumsArray()[:6], otherBall.GetBallNumsArray()[:6])
	return len(intersect) >= 4
}

func InsertBalls(ctx context.Context, balls []*Ball) error {
	err := dal.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "lottery_drawing_time"}}, // key colum
		DoUpdates: clause.AssignmentColumns([]string{
			"num1",
			"num2",
			"num3",
			"num4",
			"num5",
			"num6",
			"num7"}), // column needed to be updated
	}).Create(&balls).Error
	if err != nil {
		log.Error(ctx, "insert balls error", zap.Error(err))
		return err
	}
	return nil
}

func FindBalls(ctx context.Context, pageSize, pageNum int64) ([]*Ball, error) {
	res := make([]*Ball, 0)
	err := dal.DB.Model(&Ball{}).Offset(int(pageSize * (pageNum - 1))).Limit(int(pageSize)).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
