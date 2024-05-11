package main

import (
	"tticket/logic/tstrategy"
	"tticket/pkg/dal"
	"tticket/pkg/mail"
	"tticket/pkg/util"
	"tticket/timer"
)

func main() {
	dal.Init()

	tstrategy.Init()
	mail.Init(util.Context())
	timer.Init(util.Context())

}
