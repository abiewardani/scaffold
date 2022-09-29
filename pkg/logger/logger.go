package logger

import (
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// InitZap logger
func InitZap() {
	var (
		logg *zap.Logger
		err  error
	)

	cfg := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey: "caller",
			EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
				_, caller.File, caller.Line, _ = runtime.Caller(7)
				enc.AppendString(caller.FullPath())
			},
		},
	}

	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg.Development = true

	logg, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logg.Sync()

	// define logger
	logger = logg.Sugar()
}

// Log func
func Log(level zapcore.Level, message string, context string, scope string) {
	entry := logger.With(
		zap.String("context", context),
		zap.String("scope", scope),
	)

	switch level {
	case zapcore.DebugLevel:
		entry.Debug(message)
	case zapcore.InfoLevel:
		entry.Info(message)
	case zapcore.WarnLevel:
		entry.Warn(message)
	case zapcore.ErrorLevel:
		entry.Error(message)
	case zapcore.FatalLevel:
		entry.Fatal(message)
	case zapcore.PanicLevel:
		entry.Panic(message)
	}
}

// LogE error
func E(message interface{}) {
	logger.Error(message)
}

// LogEf error with format
func Ef(format string, i ...interface{}) {
	logger.Errorf(format, i...)
}

// LogI info
func I(message ...interface{}) {
	logger.Info(message...)
}

// LogIf info with format
func If(format string, i ...interface{}) {
	logger.Infof(format, i...)
}

// LogD info
func D(message ...interface{}) {
	logger.Debug(message...)
}

// DF info with format
func DF(format string, i ...interface{}) {
	logger.Debugf(format, i...)
}

func Panic(i ...interface{}) {
	logger.Panic(i)
}
