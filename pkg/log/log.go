package log

import (
	"context"
	"fmt"
	"os"
	"tticket/pkg/util"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	// init cfg

	file, err := os.Create("./tticket.log")
	if err != nil {
		panic(fmt.Sprintf("log file err: %v", err))
	}
	writer := zapcore.AddSync(file)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writer,
		zapcore.DebugLevel,
	)

	logger = zap.New(core, zap.AddCallerSkip(1))
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getExtraField(ctx)...)

	logger.Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getExtraField(ctx)...)

	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getExtraField(ctx)...)
	logger.Error(msg, fields...)
}

func getExtraField(ctx context.Context) []zap.Field {
	res := make([]zap.Field, 0)
	id, ok := ctx.Value(util.LOG_ID).(string)
	if id != "" && ok {
		res = append(res, zap.String(util.LOG_ID, id))
	}

	eventID, ok := ctx.Value(util.TASK_EVENT_ID).(string)
	if eventID != "" && ok {
		res = append(res, zap.String(util.TASK_EVENT_ID, eventID))
	}
	return res
}
