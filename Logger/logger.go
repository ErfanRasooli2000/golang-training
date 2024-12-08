package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {

	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "Date"
	encoderConfig.MessageKey = "Message"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig
	Log, err = config.Build()

	if err != nil {

		panic(err)

	}
}

func Info(message string, fields ...zap.Field) {

	Log.Info(message, fields...)

}

func Error(message string, fields ...zap.Field) {

	Log.Info(message, fields...)

}

func Debug(message string, fields ...zap.Field) {

	Log.Info(message, fields...)

}
