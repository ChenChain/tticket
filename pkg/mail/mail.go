package mail

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
	"time"
	"tticket/pkg/conf"
	"tticket/pkg/log"
)

var pool *email.Pool
var err error
var ch chan *email.Email

func Init(ctx context.Context) {
	pool, err = email.NewPool(
		fmt.Sprintf(conf.Config.Mail.Address+":%s", conf.Config.Mail.Port),
		4,
		smtp.PlainAuth("",
			conf.Config.Mail.UserName,
			conf.Config.Mail.Password,
			conf.Config.Mail.Host),
	)
	if err != nil {
		panic(fmt.Sprintf("init mail err:%v", err))
	}
	ch = make(chan *email.Email, 10)
	go send(ctx)
}

func send(ctx context.Context) {
	for item := range ch {
		for i := 0; i < 3; i++ {
			err := pool.Send(item, 10*time.Second)
			if err == nil {
				break
			}
			log.Warn(ctx, "send mail err", zap.Error(err),
				zap.Any("mail", item))
		}
		if err != nil {
			log.Error(ctx, "Failed to send mail", zap.Error(err),
				zap.Any("mail", item))
			err = nil
		}
	}
}

func Send(mail *email.Email) {
	ch <- mail
}
