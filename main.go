package main

import (
	"context"
	"tticket/pkg/mail"
	"tticket/timer"
)

func main() {
	ctx := context.TODO()
	mail.Init(ctx)
	timer.Init(ctx)

}
