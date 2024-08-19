package tmail

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"sort"
	"tticket/logic/tuser"
	"tticket/pkg/conf"
	"tticket/pkg/log"
	"tticket/pkg/mail"
	"tticket/pkg/model"
)

func SendSummary(ctx context.Context) error {
	balls, err := model.FindBalls(ctx, 3, 1)
	if err != nil {
		log.Error(ctx, "find balls err", zap.Error(err))
		return err
	}

	historyPredictTime := make([]string, 0)
	for _, b := range balls {
		historyPredictTime = append(historyPredictTime, b.LotteryDrawingTime)
	}

	predictAll, err := model.FindPredictBallByDrawingTime(historyPredictTime)
	if err != nil {
		log.Error(ctx, "find predict balls err", zap.Error(err))
		return err
	}

	mailStr := summaryContentTemplate(balls, predictAll)
	users := tuser.GetUserMails(ctx)
	for _, u := range users {
		// todo config in toml
		mailContent := &email.Email{
			From:    conf.Config.Mail.FromUser,
			To:      []string{u},
			Subject: "TTicket Summary Mail",
			HTML:    []byte(mailStr),
			Sender:  "TTicket Pro System",
		}
		mail.Send(mailContent)
	}
	return nil
}

func summaryContentTemplate(balls []*model.Ball, historyPredictBall []*model.PredictBall) string {
	sort.Slice(balls, func(i, j int) bool {
		return balls[i].LotteryDrawingTime < balls[j].LotteryDrawingTime
	})
	ballWinCnt := make(map[string]int)
	for _, ball := range balls {
		if _, ok := ballWinCnt[ball.LotteryDrawingTime]; !ok {
			ballWinCnt[ball.LotteryDrawingTime] = 0
		}
		for _, pre := range historyPredictBall {
			if pre.PredictLotteryDrawingTime != ball.LotteryDrawingTime {
				continue
			}
			if pre.IsWinning(ball) {
				ballWinCnt[ball.LotteryDrawingTime] = ballWinCnt[ball.LotteryDrawingTime] + 1
			}
		}
	}

	strTemplate := `
	<h3><b>日期：%s： 中奖次数：%d</b></h3><br />
`
	str := ""
	for _, ball := range balls {
		str += fmt.Sprintf(strTemplate, ball.LotteryDrawingTime, ballWinCnt[ball.LotteryDrawingTime])
	}

	return str
}
