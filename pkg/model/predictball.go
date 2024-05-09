package model

import "time"

type PredictBall struct {
	ID                        int64 `gorm:"primary_key" json:"id"`
	LotteryDrawingTime        time.Time
	PredictLotteryDrawingTime time.Time
	Num1                      int64
	Num2                      int64
	Num3                      int64
	Num4                      int64
	Num5                      int64
	Num6                      int64
	Num7                      int64
	CreatedTime               time.Time
	UpdatedTime               time.Time
	Strategy                  string
}
