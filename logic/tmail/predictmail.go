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

func Send(ctx context.Context) error {
	predictDateTime := model.GetPredictLotteryDrawingTime()
	predicts, err := model.FindPredictBallByDrawingTime([]string{predictDateTime})
	if err != nil {
		log.Error(ctx, "find predict balls err", zap.Error(err))
		return err
	}

	mailStr := contentTemplate(predicts)
	users := tuser.GetUserMails(ctx)
	for _, u := range users {
		// todo config in toml
		mailContent := &email.Email{
			From:    conf.Config.Mail.FromUser,
			To:      []string{u},
			Subject: predictDateTime + " TTicket 预测邮件",
			HTML:    []byte(mailStr),
			Sender:  "TTicket Pro System",
		}
		mail.Send(mailContent)
	}
	return nil
}

func contentTemplate(predicts []*model.PredictBall) string {
	sort.Slice(predicts, func(i, j int) bool {
		return predicts[i].OrderNum < predicts[j].OrderNum
	})

	predictStrTemplate := `
	<h3><b>%d： %s</b></h3><br />
`
	predictStr := ""
	for _, predict := range predicts {
		predictStr += fmt.Sprintf(predictStrTemplate, predict.OrderNum, predict.GetBallNumsString())
	}
	predictStr += "<br/> <br/> <br/>"

	return predictStr
}
