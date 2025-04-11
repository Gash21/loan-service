package logger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(serviceName, level string) *zap.Logger {
	syncer := zapcore.AddSync(os.Stdout)

	core := ecszap.NewCore(ecszap.NewDefaultEncoderConfig(), syncer, getLogLevel(level))
	log := zap.New(core, zap.AddCaller())
	log = log.With(zap.String("service", serviceName))

	return log
}
