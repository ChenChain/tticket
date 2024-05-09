package log

import (
	"context"
	"fmt"
	"os"

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
	logger.Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
