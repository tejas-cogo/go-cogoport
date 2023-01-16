package helpers

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func logger() {
	filename := "logs.log"
	logger := fileLogger(filename)

	logger.Info("INFO log level message")
	logger.Warn("Warn log level message")
	logger.Error("Error log level message")

}

func fileLogger(filename string) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
