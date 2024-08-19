package main

import (
	"time"
	"tticket/logic/tstrategy"
	"tticket/logic/tuser"
	"tticket/pkg/conf"
	"tticket/pkg/dal"
	"tticket/pkg/mail"
	"tticket/pkg/util"
	"tticket/timer"
)

func main() {
	conf.Init()
	dal.Init()
	if err := tuser.CacheUserMail(util.Context()); err != nil {
		panic(err)
	}
	tstrategy.Init()
	mail.Init(util.Context())
	timer.Init(util.Context())
	for {
		time.Sleep(time.Second * 1)
		// todo: ops api
	}
}
