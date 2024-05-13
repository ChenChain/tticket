package main

import (
	"tticket/logic/tstrategy"
	"tticket/pkg/conf"
	"tticket/pkg/dal"
	"tticket/pkg/mail"
	"tticket/pkg/util"
	"tticket/timer"
)

func main2() {
	conf.Init()
	dal.Init()

	tstrategy.Init()
	mail.Init(util.Context())
	timer.Init(util.Context())
	for {
		// todo: ops api
	}
}
