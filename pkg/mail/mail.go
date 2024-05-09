package mail

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
	"time"
	"tticket/pkg/log"
)

var pool *email.Pool
var err error
var ch chan *email.Email

func Init(ctx context.Context) {
	pool, err = email.NewPool(
		"smtp.126.com:25",
		4,
		smtp.PlainAuth("", "leedarjun@126.com", "358942617ldj", "smtp.126.com"),
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
