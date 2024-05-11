package tmail

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"tticket/logic/tuser"
	"tticket/pkg/log"
	"tticket/pkg/mail"
	"tticket/pkg/model"
)

type TMail struct {
}

func Send(ctx context.Context) error {
	predicts, err := model.FindPredictBalls(ctx, 4, 1)
	if err != nil {
		log.Error(ctx, "find predict balls err", zap.Error(err))
		return err
	}
	balls, err := model.FindBalls(ctx, 3, 1)
	if err != nil {
		log.Error(ctx, "find balls err", zap.Error(err))
		return err
	}
	mailStr := contentTemplate(predicts[0], balls, predicts[1:])
	users := tuser.GetUserMails(ctx)
	for _, u := range users {
		mailContent := &email.Email{
			From:    "1793854955@qq.com",
			To:      []string{u},
			Subject: "TTicket 预测邮件",
			HTML:    []byte(mailStr),
			Sender:  "TTicket System",
		}
		mail.Send(mailContent)
	}
	return nil
}

func contentTemplate(predict *model.PredictBall, balls []*model.Ball, historyPredictBall []*model.PredictBall) string {
	predictStr := `
	<h3><b>预测日期/期号：%s</b></h3>
	<h3><b>结果：%s</b></h3>
`
	predictStr = fmt.Sprintf(predictStr, predict.PredictLotteryDrawingTime, predict.GetBallNumsString())

	predictBallMap := make(map[string]*model.PredictBall)
	for _, v := range historyPredictBall {
		predictBallMap[v.PredictLotteryDrawingTime] = v
	}
	for _, v := range balls {
		if predictball, ok := predictBallMap[v.LotteryDrawingTime]; ok {
			if predictball.IsWinning(v) {
				v.LotteryDrawingTime = "<b> 中奖：" + v.LotteryDrawingTime + " </b>"
			}
		} else {
			// 防止下面空指针，塞一个异常值
			predictBallMap[v.LotteryDrawingTime] = model.GenerateErrorPredictBall()
		}
	}

	historyStr := `
<table>
  <thead>
    <tr>
      <th>日期/期号</th>
      <th>实际</th>
      <th>预测</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>%s</td>
      <td>%s</td>
      <td>%s</td>
    </tr>
    <tr>
      <td>%s</td>
      <td>%s</td>
      <td>%s</td>
    </tr>
    <tr>
      <td>%s</td>
      <td>%s</td>
      <td>%s</td>
    </tr>
  </tbody>
</table>
`
	historyStr = fmt.Sprintf(historyStr,
		balls[0].LotteryDrawingTime, balls[0].GetBallNumsString(), predictBallMap[balls[0].LotteryDrawingTime],
		balls[1].LotteryDrawingTime, balls[1].GetBallNumsString(), predictBallMap[balls[1].LotteryDrawingTime],
		balls[2].LotteryDrawingTime, balls[2].GetBallNumsString(), predictBallMap[balls[2].LotteryDrawingTime],
	)
	return predictStr + historyStr
}
