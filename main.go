package main

import (
	"context"
	"tticket/pkg/mail"
)

func main() {
	mail.Init(context.Background())
}
