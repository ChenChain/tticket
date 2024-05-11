package util

import (
	"context"
	"github.com/google/uuid"
)

var LOG_ID = "log_id"
var TASK_EVENT_ID = "event_id"

func Context() context.Context {
	ctx := context.TODO()
	return context.WithValue(ctx, LOG_ID, uuid.New().String())
}
