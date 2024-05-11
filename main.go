package main

import (
	"tticket/pkg/dal"
	"tticket/pkg/mail"
	"tticket/pkg/util"
	"tticket/timer"
)

func main() {
	dal.Init()
	mail.Init(util.Context())
	timer.Init(util.Context())

}
