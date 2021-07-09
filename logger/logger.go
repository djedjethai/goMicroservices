package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // zapcore is undefined ??
	config.EncoderConfig = encoderConfig

	// with the config, instead of the following line
	log, err = config.Build(zap.AddCallerSkip(1))

	// here only one skip level so one, level?? maybe nested docs..?
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
